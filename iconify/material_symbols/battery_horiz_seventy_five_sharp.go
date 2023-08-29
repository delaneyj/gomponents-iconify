package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryHorizSeventyFiveSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17v-3H2v-4h2V7h18v10H4Zm2-2h4V9H6v6Z"/>`),
		g.Group(children),
	)
}