package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneKPlusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 14h1v-1.5H19v-1h-1.5V10h-1v1.5H15v1h1.5V14Zm-6 1H12v-2.25L13.75 15h1.75l-2.25-3l2.25-3h-1.75L12 11.25V9h-1.5v6Zm-3 0H9V9H6v1.5h1.5V15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}