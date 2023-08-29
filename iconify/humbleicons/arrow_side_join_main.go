package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSideJoinMain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15h13.01m0 0a6 6 0 0 1-5.23-3.058l-1.06-1.884A6 6 0 0 0 4.49 7H3m13.01 8H21m0 0l-3 3m3-3l-3-3"/>`),
		g.Group(children),
	)
}