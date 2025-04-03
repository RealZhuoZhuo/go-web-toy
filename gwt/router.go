package gwt

import "strings"

type Router struct {
	handlers map[string]Handlerfunc
}

func parsePath(path string) (parts []string) {
	tem := strings.Split(path, "/")
	for _, p := range tem {
		if p != "" {
			parts = append(parts, p)
			if p == "*" {
				break
			}
		}
	}
	return
}
func newRouter() *Router {
	return &Router{make(map[string]Handlerfunc)}
}
func (r *Router) addRoute(method string, pattern string, handler Handlerfunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}
