package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectMoveDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h18v2H3Zm9 12l-4-4l1.4-1.4l1.6 1.575V7h2v6.175l1.6-1.575L16 13l-4 4Zm-9 4v-2h18v2H3Z"/>`),
		g.Group(children),
	)
}