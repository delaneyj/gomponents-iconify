package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 18a2 2 0 1 0 4 0a2 2 0 1 0-4 0M17 6a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M19 8v5c0 .594-.104 1.164-.294 1.692m-1.692 2.298A4.978 4.978 0 0 1 14 18h-3l3-3m0 6l-3-3m-6-2v-5c0-1.632.782-3.082 1.992-4M10 6h3l-3-3m1.501 4.499L13 6M3 3l18 18"/></g>`),
		g.Group(children),
	)
}