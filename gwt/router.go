package gwt

import (
	"net/http"
	"strings"
)

type Router struct {
	handlers map[string]Handlerfunc
	root     map[string]*node
}
type node struct {
	children map[string]*node
	isWild   bool
	path     string
	part     string
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
	return &Router{
		make(map[string]Handlerfunc),
		make(map[string]*node),
	}
}
func (r *Router) addRoute(method string, path string, handler Handlerfunc) {
	key := method + "-" + path
	r.handlers[key] = handler
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{
			children: make(map[string]*node),
		}
	}
	root := r.root[method]
	parts := parsePath(path)
	for _, part := range parts {
		if _, ok := root.children[part]; !ok {
			root.children[part] = &node{
				part:     part,
				children: make(map[string]*node),
				isWild:   part[0] == ':' || part[0] == '*',
			}
		}
		root = root.children[part]
	}
	root.path = path
}
func (r *Router) getRoute(method, path string) (*node, map[string]string) {
	params := map[string]string{}
	searchParts := parsePath(path)
	if _, ok := r.root[method]; !ok {
		return nil, nil
	}
	root := r.root[method]
	for i, part := range searchParts {
		tem := ""
		for _, child := range root.children {
			if child.part == part || child.isWild {
				if child.part[0] == ':' {
					params[child.part[1:]] = part
				}
				if child.part[0] == '*' {
					params[child.part[1:]] = strings.Join(searchParts[i:], "/")
				}
				tem = child.part
			}
		}
		if tem == "" {
			return nil, nil
		}
		if tem[0] == '*' {
			return root.children[tem], params
		}
		root = root.children[tem]
	}
	return root, params
}
func (r *Router) Handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.path
		if handler, ok := r.handlers[key]; ok {
			c.handlers = append(c.handlers, handler)
		} else {
			http.NotFound(c.W, c.R)
		}
	} else {
		http.NotFound(c.W, c.R)
	}
	c.Next()
}
