package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h18v2H3Zm6 4V7h12v2H9Zm-6 4v-2h18v2H3Zm6 4v-2h12v2H9Zm-6 4v-2h18v2H3Z"/>`),
		g.Group(children),
	)
}