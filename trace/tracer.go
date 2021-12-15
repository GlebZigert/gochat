package trace


import(
"io"
"fmt"

)

type tracer struct {
	out io.Writer
	}
	func (t *tracer) Trace(a ...interface{}) {

		
		t.out.Write([]byte(fmt.Sprint(a...)))
		t.out.Write([]byte("\n"))

		}
		
	


// Tracer is the interface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}

}

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}


	
