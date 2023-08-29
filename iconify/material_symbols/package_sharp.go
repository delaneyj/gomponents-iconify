package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h5v-2H7v2Zm-4 4V3h18v18H3ZM8 5v8l4-2l4 2V5H8Z"/>`),
		g.Group(children),
	)
}