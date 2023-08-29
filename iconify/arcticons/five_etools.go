package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveEtools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.31 33.03c1.5 1.257 2.955 1.833 6.592 1.833h.784a7.018 7.018 0 0 0 7.018-7.018h0a7.018 7.018 0 0 0-7.018-7.018H11.31v-7.69h13.566m11.27 14.272a4.154 4.154 0 0 1-3.61 2.096h0a4.155 4.155 0 0 1-4.155-4.155v-2.7a4.155 4.155 0 0 1 4.155-4.155h0a4.155 4.155 0 0 1 4.154 4.155V24h-8.31"/>`),
		g.Group(children),
	)
}