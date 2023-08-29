package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airpods(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M14 7c0 2.21 1.644 4 4 4s4-1.79 4-4s-1.644-4-4-4s-4 1.79-4 4Zm-4 0c0 2.21-1.644 4-4 4S2 9.21 2 7s1.644-4 4-4s4 1.79 4 4Z"/><path stroke-linecap="round" d="M14 7v12a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2v-8"/><path d="M14 17h4M6 17h4"/><path stroke-linecap="round" d="M10 7v12a2 2 0 0 1-2 2v0a2 2 0 0 1-2-2v-8"/><path d="M20 4a5.408 5.408 0 0 0 0 6M4 4a5.408 5.408 0 0 1 0 6"/></g>`),
		g.Group(children),
	)
}