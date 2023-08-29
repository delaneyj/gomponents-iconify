package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanToolAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.525 21L1.15 11.925l1.775-1.7L7 13.075V2h2v9h2V6h2v5h2V7h2v4h2V9h2v12H8.525Z"/>`),
		g.Group(children),
	)
}