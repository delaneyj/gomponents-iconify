package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TenKSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h1.5V9H5v1.5h1V15Zm2.5 0H13V9H8.5v6Zm1.5-1.5v-3h1.5v3H10Zm4 1.5h1.5v-2.25L17.25 15H19l-2.25-3L19 9h-1.75l-1.75 2.25V9H14v6ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}