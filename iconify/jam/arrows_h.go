package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.423 4.996h13.154L14.04 2.46a1 1 0 1 1 1.415-1.414l4.242 4.243a1 1 0 0 1 0 1.414l-4.242 4.242a1 1 0 0 1-1.415-1.414l2.536-2.535H3.423L5.96 9.53a1 1 0 1 1-1.415 1.414L.302 6.703a.997.997 0 0 1 0-1.414l4.242-4.243A1 1 0 1 1 5.96 2.46L3.423 4.996z"/>`),
		g.Group(children),
	)
}