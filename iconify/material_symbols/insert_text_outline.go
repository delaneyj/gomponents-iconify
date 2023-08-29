package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertTextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16v-6H8V8h8v2h-3v6h-2ZM1 23v-6h2V7H1V1h6v2h10V1h6v6h-2v10h2v6h-6v-2H7v2H1Zm6-4h10v-2h2V7h-2V5H7v2H5v10h2v2ZM3 5h2V3H3v2Zm16 0h2V3h-2v2Zm0 16h2v-2h-2v2ZM3 21h2v-2H3v2ZM5 5Zm14 0Zm0 14ZM5 19Z"/>`),
		g.Group(children),
	)
}