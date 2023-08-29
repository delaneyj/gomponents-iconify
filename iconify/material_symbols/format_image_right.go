package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatImageRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17V7h10v10H11Zm2-2h6V9h-6v6ZM3 21v-2h18v2H3Zm0-4v-2h6v2H3Zm0-4v-2h6v2H3Zm0-4V7h6v2H3Zm0-4V3h18v2H3Z"/>`),
		g.Group(children),
	)
}