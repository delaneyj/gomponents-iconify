package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewWindowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h8v2H5v14h14v-6h2v8H3Zm13-10V8h-3V6h3V3h2v3h3v2h-3v3h-2Z"/>`),
		g.Group(children),
	)
}