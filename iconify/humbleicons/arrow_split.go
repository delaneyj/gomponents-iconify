package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h4.597a5 5 0 0 1 3.904 1.877l.998 1.246A5 5 0 0 0 16.403 17H21m0 0l-3-3m3 3l-3 3m3-13h-5.078A4 4 0 0 0 12.8 8.501L11.201 10.5A4 4 0 0 1 8.078 12H6m15-5l-3-3m3 3l-3 3"/>`),
		g.Group(children),
	)
}