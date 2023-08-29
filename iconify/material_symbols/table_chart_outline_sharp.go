package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChartOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h19v18H3ZM5 8h15V5H5v3Zm3 2H5v9h3v-9Zm9 0v9h3v-9h-3Zm-2 0h-5v9h5v-9Z"/>`),
		g.Group(children),
	)
}