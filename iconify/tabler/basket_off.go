package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 10l-2-6m-8 6l.75-2.252m1.001-3.002L9 4m3 4h7a2 2 0 0 1 1.977 2.304C20.535 12.82 20.221 14.742 20 16m-1.01 3.003a2.997 2.997 0 0 1-2.234.997H7.244a3 3 0 0 1-2.965-2.544l-1.255-7.152A2 2 0 0 1 5.001 8H8"/><path d="M12 12a2 2 0 1 0 2 2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}