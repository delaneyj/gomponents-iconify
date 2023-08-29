package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16v-6H8V8h8v2h-3v6h-2ZM1 23v-6h2V7H1V1h6v2h10V1h6v6h-2v10h2v6h-6v-2H7v2H1Zm6-4h10v-2h2V7h-2V5H7v2H5v10h2v2Z"/>`),
		g.Group(children),
	)
}