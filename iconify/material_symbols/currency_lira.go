package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyLira(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-5.15l-3 1.875V15.35l3-1.875v-2.35L6 13v-2.35l3-1.9V3h2v4.5L15 5v2.35l-4 2.525v2.35l4-2.5v2.35L11 14.6V19q2.075 0 3.538-1.463T16 14h2q0 2.925-2.038 4.963T11 21H9Z"/>`),
		g.Group(children),
	)
}