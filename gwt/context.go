package gwt

import "net/http"

type Context struct {
	W          http.ResponseWriter
	R          *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
	// middlewares
	handlers []Handlerfunc
	index    int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		R:      r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}
func (c *Context) Next() {
	c.index++ //如果有next的话
	s := len(c.handlers)
	for ; c.index < s; c.index++ { //没有next的话
		c.handlers[c.index](c)
	}
}
