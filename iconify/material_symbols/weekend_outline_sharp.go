package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeekendOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V9h3V4h16v5h3v11H1Zm6-6h10V9h1V6H6v3h1v5Zm-4 4h18v-7h-2v5H5v-5H3v7Zm9-2Zm0-2Zm0 2Z"/>`),
		g.Group(children),
	)
}