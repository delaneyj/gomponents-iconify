package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.147 21.197l1.67-1.168a13.393 13.393 0 0 0 5.447-13.624a.837.837 0 0 0-.453-.586L12 2L4.19 5.819a.837.837 0 0 0-.454.586a13.393 13.393 0 0 0 5.448 13.624l1.67 1.168a2 2 0 0 0 2.293 0Z"/>`),
		g.Group(children),
	)
}