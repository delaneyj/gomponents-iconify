package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkFiveGNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsNetwork5gNegative0)"><path d="M22 30a1 1 0 0 0-1-1h-5.25a1 1 0 0 0-.98.804l-.75 3.75a1 1 0 0 0 1.535 1.028c.21-.14 1.143-.582 2.445-.582a2 2 0 1 1 0 4h-.832c-.577 0-1.054-.36-1.227-.84a1 1 0 1 0-1.882.68A3.297 3.297 0 0 0 17.17 40H18a4 4 0 0 0 0-8a7.72 7.72 0 0 0-1.666.182L16.57 31H21a1 1 0 0 0 1-1Zm10.753 1.774a1 1 0 0 0-.08-1.412a5.503 5.503 0 1 0 0 8.204a1 1 0 0 0 .334-.745v-3.357a1 1 0 0 0-1-1h-3.002a1 1 0 1 0 0 2h2.002v1.875a3.504 3.504 0 1 1 .334-5.486a1 1 0 0 0 1.412-.079Z"/><path fill-rule="evenodd" d="M0 0h48v48H0V0Zm10 29a3 3 0 0 1 3-3h22a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H13a3 3 0 0 1-3-3V29ZM32.485 5.515a12.001 12.001 0 0 1 0 16.97L33.9 23.9a14.001 14.001 0 0 0 0-19.8l-1.415 1.415Zm-2.942 6.189a6.001 6.001 0 0 0-1.3-1.947l1.414-1.414a8 8 0 0 1 0 11.314l-1.414-1.414a6 6 0 0 0 1.3-6.54Zm-11.086 0a6.001 6.001 0 0 1 1.3-1.947l-1.414-1.414a8 8 0 0 0 0 11.314l1.414-1.414a6 6 0 0 1-1.3-6.54Zm-2.942-6.19a12.001 12.001 0 0 0 0 16.971L14.1 23.9a14.001 14.001 0 0 1 0-19.8l1.415 1.415ZM24 16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsNetwork5gNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}