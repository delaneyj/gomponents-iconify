package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19h2v-2H7v2Zm0-4h2v-4H7v4Zm4 4h2v-4h-2v4Zm0-6h2v-2h-2v2Zm4 6h2v-2h-2v2Zm0-4h2v-4h-2v4ZM4 22V8l6-6h10v20H4Z"/>`),
		g.Group(children),
	)
}