package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AncientPavilionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.516 7.877A4.5 4.5 0 0 0 5 11.887v7.114H3v2h18v-2h-2v-7.114a4.5 4.5 0 0 0 3.484-4.01A9 9 0 0 1 12.514 2h-1.029a9 9 0 0 1-9.97 5.877ZM17 19H7v-7h10v7.001Z"/>`),
		g.Group(children),
	)
}