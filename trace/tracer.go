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


	
