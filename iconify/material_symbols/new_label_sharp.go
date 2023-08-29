package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewLabelSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-3H2v-2h3v-3h2v3h3v2H7v3H5Zm7-1v-6H9v-3H3V5h13.05L21 12l-4.95 7H12Z"/>`),
		g.Group(children),
	)
}