package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoKPlusSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 14h1v-1.5H19v-1h-1.5V10h-1v1.5H15v1h1.5V14ZM11 15h1.5v-2.25L14.25 15H16l-2.25-3L16 9h-1.75l-1.75 2.25V9H11v6Zm-5 0h4v-1.5H7.5v-1H10V9H6v1.5h2.5v1H6V15Zm-3 6V3h18v18H3Z"/>`),
		g.Group(children),
	)
}