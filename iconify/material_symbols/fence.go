package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-4H3v-2h2v-2H3v-2h2V7l3-3l2 2l2.025-2l2 2l2-2l3 3v3H21v2h-1.975v2H21v2h-1.975v4H5Zm2-10h2V7.825l-1-1l-1 1V10Zm4 0h2V7.825l-1-1l-1 1V10Zm4.025 0H17V7.825l-1-1l-.975.975V10ZM7 14h2v-2H7v2Zm4 0h2v-2h-2v2Zm4.025 0H17v-2h-1.975v2ZM7 18h2v-2H7v2Zm4 0h2v-2h-2v2Zm4.025 0H17v-2h-1.975v2Z"/>`),
		g.Group(children),
	)
}