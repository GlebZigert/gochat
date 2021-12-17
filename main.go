package main

import (
	"flag"
	"fmt"
	"gochat/trace"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"	
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
		}
		if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
		}
		fmt.Println("data[UserData]", data["UserData"])
		


	t.templ.Execute(w, data)
}

func main() {

	fmt.Println(trace.ReverseRunes(".olleH"))

	var addr = flag.String("addr", ":8080", "the addr of the application.")

	flag.Parse()

	// set up gomniauth
	gomniauth.SetSecurityKey("some long key")
	gomniauth.WithProviders(
		facebook.New("key", "secret",
			"http://localhost:8080/auth/callback/facebook"),
		github.New("key", "secret",
			"http://localhost:8080/auth/callback/github"),
		google.New("234438523275-p23ad7cql6c66mpdsq0d2097ei9smv7d.apps.googleusercontent.com", 
		"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",
			"http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	//	r.tracer = trace.New(os.Stdout)
	//	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	http.Handle("/room", r)
	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
