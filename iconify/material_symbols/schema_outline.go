package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchemaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23v-6h2.5v-2H4V9h2.5V7H4V1h7v6H8.5v2H11v2h3V9h7v6h-7v-2h-3v2H8.5v2H11v6H4Zm2-2h3v-2H6v2Zm0-8h3v-2H6v2Zm10 0h3v-2h-3v2ZM6 5h3V3H6v2Zm1.5-1Zm0 8Zm10 0Zm-10 8Z"/>`),
		g.Group(children),
	)
}