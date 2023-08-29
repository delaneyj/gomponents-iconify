package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 21h9m-4 0l-7-8l8.5-5.5"/><path d="M13 14c-2.148-2.148-2.148-5.852 0-8c2.088-2.088 5.842-1.972 8 0l-8 8z"/><path d="m11.742 7.574l-1.156-1.156a2 2 0 0 1 2.828-2.829l1.144 1.144M15.5 12l.208.274a2.527 2.527 0 0 0 3.556 0c.939-.933.98-2.42.122-3.4l-.366-.369"/></g>`),
		g.Group(children),
	)
}