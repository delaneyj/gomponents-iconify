package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatImageLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7h10v10H3Zm2-2h6V9H5v6ZM3 5V3h18v2H3Zm12 4V7h6v2h-6Zm0 4v-2h6v2h-6Zm0 4v-2h6v2h-6ZM3 21v-2h18v2H3Z"/>`),
		g.Group(children),
	)
}