package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeCells(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M20 14V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v38a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9m8 0v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H29a1 1 0 0 0-1 1v9m0 10h16M5 24h15"/><path stroke-linejoin="round" d="m32.748 28.818l-1.59-1.59l-3.182-3.183l3.181-3.182l1.591-1.59m-17.373 9.545l1.591-1.59l3.182-3.183l-3.182-3.182l-1.591-1.59"/></g>`),
		g.Group(children),
	)
}