package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAdOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h16V6H4v12Zm0 0V6v12Zm1-1h14v-2H5v2Zm0-4h14v-2H5v2Zm0-4h10V7H5v2Z"/>`),
		g.Group(children),
	)
}