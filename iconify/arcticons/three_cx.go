package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeCx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M30.295 17.807L38.5 30.193m0-12.386l-8.205 12.386m-1.735-3.686a4.103 4.103 0 0 1-4.082 3.686h0a4.103 4.103 0 0 1-4.102-4.103v-4.18a4.103 4.103 0 0 1 4.102-4.102h0a4.103 4.103 0 0 1 4.082 3.683M9.501 29.148c.855.716 1.779 1.045 3.852 1.045h1.257a3.096 3.096 0 0 0 3.095-3.096h0a3.096 3.096 0 0 0-3.095-3.096M9.5 18.842c.857-.714 1.78-1.04 3.854-1.035l1.256.003a3.096 3.096 0 0 1 3.095 3.096h0A3.096 3.096 0 0 1 14.61 24m-3.154.001h3.154"/><path d="m28.581 24.001l-1.926-1.882v3.765l1.926-1.883z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}