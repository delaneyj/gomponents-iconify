package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DescriptionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18h8v-2H8v2Zm0-4h8v-2H8v2Zm-4 8V2h10l6 6v14H4Zm9-13V4H6v16h12V9h-5ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}