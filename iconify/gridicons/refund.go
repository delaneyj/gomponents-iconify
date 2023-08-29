package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refund(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.91 2.91L11.83 5H14a8 8 0 0 1 8 8h-2a6 6 0 0 0-6-6h-2.17l2.09 2.09l-1.42 1.41L8 6l1.41-1.41L12.5 1.5l1.41 1.41zM2 12v10h16V12H2zm2 6.56v-3.11A4 4 0 0 0 5.45 14h9.1A4 4 0 0 0 16 15.45v3.11A3.996 3.996 0 0 0 14.57 20H5.45A3.996 3.996 0 0 0 4 18.56zm6 .44c.828 0 1.5-.895 1.5-2s-.672-2-1.5-2s-1.5.895-1.5 2s.672 2 1.5 2z"/>`),
		g.Group(children),
	)
}