package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelImportantOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 19l4.5-7L3 5h13.05L21 12l-4.95 7H3Zm3.65-2H15l3.55-5L15 7H6.65l3.25 5l-3.25 5Zm3.25-5L6.65 7l3.25 5l-3.25 5l3.25-5Z"/>`),
		g.Group(children),
	)
}