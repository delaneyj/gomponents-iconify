package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 8v8l5.333-4L14 8zm-2-1c0-1.236 1.411-1.942 2.4-1.2l6.667 5c.8.6.8 1.8 0 2.4l-6.667 5c-.989.742-2.4.036-2.4-1.2v-3l-5.6 4.2c-.989.742-2.4.036-2.4-1.2V7c0-1.236 1.411-1.942 2.4-1.2L12 10V7zm-.667 5L6 8v8l5.333-4z"/></g>`),
		g.Group(children),
	)
}