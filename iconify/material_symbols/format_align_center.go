package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm4-4v-2h10v2H7Zm-4-4v-2h18v2H3Zm4-4V7h10v2H7ZM3 5V3h18v2H3Z"/>`),
		g.Group(children),
	)
}