package router

import (
	"strings"
	"syscall/js"
)

type Router struct {
	routes map[string]func(params map[string]string)
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]func(params map[string]string)),
	}
}

func (r *Router) Register(path string, handler func(params map[string]string)) {
	r.routes[path] = handler
}

func (r *Router) Navigate(path string) {
	cleanPath := strings.TrimPrefix(path, "/")

	if handler, exists := r.routes[cleanPath]; exists {
		handler(nil)
		return
	}

	if handler, exists := r.routes["404"]; exists {
		handler(nil)
	}
}

func (r *Router) SetupEventListeners() {

	// Handle Links
	js.Global().Get("document").Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		e := args[0]
		target := e.Call("target").Call("closest", "a")
		if !target.IsNull() {
			href := target.Get("pathname").String()
			if strings.HasPrefix(href, "/") {
				e.Call("preventDefault")
				js.Global().Get("history").Call("pushState", nil, "", href)
				r.Navigate(href)
			}
		}
		return nil
	}))

	// Handle Pops
	js.Global().Get("window").Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		path := js.Global().Get("location").Get("pathname").String()
		r.Navigate(path)
		return nil
	}))

	// Init
	path := js.Global().Get("location").Get("pathname").String()
	r.Navigate(path)
}
