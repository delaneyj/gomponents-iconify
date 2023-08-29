package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabProfileOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 12v-2h8v2H8Zm0-4V6h8v2H8Zm-2 6h8.975L18 17.95V4H6v10Zm0 6h11.05L14 16H6v4Zm14 2H4V2h16v20ZM6 20V4v16Zm0-4v-2v2Z"/>`),
		g.Group(children),
	)
}