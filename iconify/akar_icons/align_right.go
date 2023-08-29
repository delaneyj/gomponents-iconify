package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M22 3v18"/><path stroke-linejoin="round" d="M2 6h16v4H2zm6 8h10v4H8z"/></g>`),
		g.Group(children),
	)
}