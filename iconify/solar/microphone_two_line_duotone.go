package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 8a5 5 0 0 1 10 0v3a5 5 0 0 1-10 0V8Z"/><path stroke-linecap="round" d="M13.5 8H17m-3.5 3H17M7 8h2m-2 3h2m11-1v1a8 8 0 0 1-8 8m-8-9v1a8 8 0 0 0 8 8m0 0v3" opacity=".5"/></g>`),
		g.Group(children),
	)
}