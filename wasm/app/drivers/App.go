package drivers

import (
	"syscall/js"
	"wasm/app/api"
	"wasm/app/router"
	"wasm/app/templates"

	"github.com/brightsidedeveloper/goat"
)

func RouterProvider() {

	router := router.NewRouter()

	router.Register("", func(params map[string]string) {
		goat.RenderRoot(App())
	})

	router.Register("page", func(params map[string]string) {
		goat.RenderRoot(templates.Page())
	})

	router.Register("404", func(params map[string]string) {
		goat.RenderRoot(templates.NotFound())
	})

	router.SetupEventListeners()

}

var message string

func App() goat.TemplJoint {

	msg, err := api.GetWasm()

	if msg != "" {
		goat.Log("Also, the server wanted you to know: ", msg)
	}

	goat.JSFunc("goat", func(this js.Value, args []js.Value) any {
		message = "FACTS"
		goat.Render("App", templates.App(message, nil))
		return nil
	})

	message = msg

	return templates.App(message, err)
}
