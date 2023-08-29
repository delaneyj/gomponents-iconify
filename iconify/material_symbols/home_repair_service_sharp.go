package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRepairServiceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8h6V6H9v2ZM2 20v-5h4v1h2v-1h8v1h2v-1h4v5H2Zm0-6V8h5V4h10v4h5v6h-4v-2h-2v2H8v-2H6v2H2Z"/>`),
		g.Group(children),
	)
}