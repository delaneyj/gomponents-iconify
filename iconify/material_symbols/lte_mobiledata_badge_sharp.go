package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LteMobiledataBadgeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16h5v-2H6V8H4v8Zm6 0h2v-6h2V8H8v2h2v6Zm5 0h5v-2h-3v-1h2v-2h-2v-1h3V8h-5v8ZM1 21V3h22v18H1Z"/>`),
		g.Group(children),
	)
}