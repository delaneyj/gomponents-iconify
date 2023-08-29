package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WordamentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2Zm10.342 0v37m12.368-37v37M5.5 17.842h37M5.5 30.21h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.758 7.764l2 8l2-8l2 8l2-8m8.242 8a2.65 2.65 0 0 1-2.65-2.65v-2.7a2.65 2.65 0 1 1 5.3 0v2.7a2.65 2.65 0 0 1-2.65 2.65Zm9.634 0v-8h2.619c1.48 0 2.681 1.203 2.681 2.686s-1.2 2.687-2.681 2.687h-2.619m2.62.001l2.62 2.624M33.64 28.029v-8h1.8a3.5 3.5 0 0 1 3.5 3.5v1a3.5 3.5 0 0 1-3.5 3.5h-1.8Zm4.379 9.584h-3.466m-.863 2.652l2.6-7.976l2.6 8M20 40.28v-7.99l4 8l4-7.989v7.988m-18.236-4h2.608m1.392 4h-4v-8h4M9.108 28v-8l5.3 8v-8m6.942 0h5.3M24 28v-8"/>`),
		g.Group(children),
	)
}