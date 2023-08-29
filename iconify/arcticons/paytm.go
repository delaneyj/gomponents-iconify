package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paytm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1m0 12.9h37m-37-3.033h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.557 23.533v-7.991l4 8l4-7.988v7.988m-31.424 0v-8h2.619a2.687 2.687 0 0 1 0 5.374h-2.62m16.516-5.374h5.3m-2.65 8v-8m-2.65 6.216v2.7a2 2 0 0 1-2 2h0a1.994 1.994 0 0 1-1.414-.586"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.648 18.458v3.3a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-3.3m-1.906 3.064a2 2 0 0 1-2 2h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2m0 3.3v-5.3"/>`),
		g.Group(children),
	)
}