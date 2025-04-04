package gwt

import "net/http"

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}
