package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15h2v-3h3.5v2.5L17 11l-3.5-3.5V10H8v5Zm4 7.8L1.2 12L12 1.2L22.8 12L12 22.8Z"/>`),
		g.Group(children),
	)
}