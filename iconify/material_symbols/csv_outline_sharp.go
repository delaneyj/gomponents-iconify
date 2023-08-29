package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CsvOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.75 15h4v-1.5h-2.5v-3h2.5V9h-4v6Zm4.9 0h4v-3.65h-2.5v-.85h2.5V9h-4v3.6h2.5v.9h-2.5V15Zm6.6 0h1.5l1.75-6H18l-1 3.45L16 9h-1.5l1.75 6ZM2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}