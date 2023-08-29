package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CashBanknoteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.88 9.878a3 3 0 1 0 4.242 4.243m.58-3.425a3.012 3.012 0 0 0-1.412-1.405"/><path d="M10 6h9a2 2 0 0 1 2 2v8c0 .294-.064.574-.178.825M18 18H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1m12 6h.01M6 12h.01M3 3l18 18"/></g>`),
		g.Group(children),
	)
}