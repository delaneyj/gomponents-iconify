package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoHistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 10.55a8 8 0 1 1 1.073 3.952"/><path fill="currentColor" fill-rule="nonzero" d="m2.5 13.5l2.5-3H0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.5v5h3"/></g>`),
		g.Group(children),
	)
}