package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnalyticsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h18v18H3Zm2-2h14V5H5v14Zm2-2h2v-5H7v5Zm8 0h2V7h-2v10Zm-4 0h2v-3h-2v3Zm0-5h2v-2h-2v2Zm-6 7V5v14Z"/>`),
		g.Group(children),
	)
}