package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M37.845 35.022L44 41.316V44H4v-2.684l6.118-6.257l27.727-.037Z"/><path stroke-linecap="round" d="M10.104 35.074L18 27V6h12v21l7.873 8.05M11 35h26m-7-21h-6m6 6h-6"/></g>`),
		g.Group(children),
	)
}