package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-9h3v-2H5V3h14v7h-3v2h3v9h-2v-3H7v3H5ZM7 8h10V5H7v3Zm3 4h4v-2h-4v2Zm-3 4h10v-2H7v2Zm0-8V5v3Zm0 8v-2v2Z"/>`),
		g.Group(children),
	)
}