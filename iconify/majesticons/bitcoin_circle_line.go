package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="M8 8h5m1 4c.667 0 2-.4 2-2s-1.333-2-2-2h-1m1 4c.667 0 2 .4 2 2s-1.333 2-2 2h-1m1-4h-4m-2 4h5M10 6v6m0 6v-6m3-6v2m0 8v2"/></g>`),
		g.Group(children),
	)
}