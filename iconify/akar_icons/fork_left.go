package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 3v18m0-18l-4 4m4-4l4 4M8 10l-4 4l4 4"/><path d="M15 21a7 7 0 0 0-7-7H4"/></g>`),
		g.Group(children),
	)
}