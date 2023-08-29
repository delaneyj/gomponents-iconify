package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.255 29.981V18.02h3.916a4.018 4.018 0 0 1 0 8.035h-3.916M9.5 29.981V18.019l7.925 11.962V18.019M30.812 28.67a3.346 3.346 0 0 0 2.933 1.311h1.77a2.987 2.987 0 0 0 2.985-2.99h0A2.987 2.987 0 0 0 35.516 24h-1.957a2.987 2.987 0 0 1-2.984-2.99h0a2.987 2.987 0 0 1 2.984-2.991h1.771a3.346 3.346 0 0 1 2.933 1.31"/>`),
		g.Group(children),
	)
}