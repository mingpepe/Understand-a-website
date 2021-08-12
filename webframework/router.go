package webframework

import "net/http"

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) add(method, pattern string, handler HandlerFunc) {
	key := toKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := toKey(c.Method, c.Path)
	var handler HandlerFunc
	if h, ok := r.handlers[key]; ok {
		handler = h
	} else {
		handler = func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		}
	}
	c.handlers = append(c.handlers, handler)
	c.Next()
}

func toKey(method, pattern string) string {
	return method + "-" + pattern
}
