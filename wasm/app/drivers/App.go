package drivers

import (
	"wasm/app/api"
	"wasm/app/templates"

	"github.com/brightsidedeveloper/goat"
)

func App() goat.TemplJoint {

	msg, err := api.GetWasm()

	if msg != "" {
		goat.Log("Also, the server wanted you to know: ", msg)
	}

	return templates.App(msg, err)
}

// You could also make a chan
// fetch in a go routine
// rerender with the data
// boom, loading state
