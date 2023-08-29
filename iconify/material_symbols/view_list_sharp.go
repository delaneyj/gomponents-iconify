package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20h13v-4H9v4ZM2 8h5V4H2v4Zm0 6h5v-4H2v4Zm0 6h5v-4H2v4Zm7-6h13v-4H9v4Zm0-6h13V4H9v4Z"/>`),
		g.Group(children),
	)
}