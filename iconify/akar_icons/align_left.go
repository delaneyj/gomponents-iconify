package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M2 3v18"/><path stroke-linejoin="round" d="M6 6h16v4H6zm0 8h10v4H6z"/></g>`),
		g.Group(children),
	)
}