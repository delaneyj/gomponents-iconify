package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 3v18M9 3l4 4M9 3L5 7m11 3l4 4l-4 4"/><path d="M9 21a7 7 0 0 1 7-7h4"/></g>`),
		g.Group(children),
	)
}