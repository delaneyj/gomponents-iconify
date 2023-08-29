package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bunq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.49 20.57v4.27a2.59 2.59 0 0 0 2.59 2.59h0a2.58 2.58 0 0 0 2.58-2.59v-4.27m7.93 6.86v-4.27A2.59 2.59 0 0 0 28 20.57h0a2.58 2.58 0 0 0-2.58 2.59v4.27M9.5 23.16a2.59 2.59 0 0 1 2.59-2.59h0a2.59 2.59 0 0 1 2.59 2.59v1.68a2.59 2.59 0 0 1-2.59 2.59h0a2.59 2.59 0 0 1-2.59-2.59m0 2.59V17.08m29 7.76a2.59 2.59 0 0 1-2.59 2.59h0a2.59 2.59 0 0 1-2.59-2.59v-1.68a2.59 2.59 0 0 1 2.59-2.59h0a2.59 2.59 0 0 1 2.59 2.59m0-2.59v10.35"/>`),
		g.Group(children),
	)
}