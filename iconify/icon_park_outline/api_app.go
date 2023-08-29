package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApiApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M18 23.937V10a6 6 0 0 1 12 0v2.006m0 11.997V38a6 6 0 0 1-12 0v-2.03"/><path d="M24 30H9.984C6.68 30 4 27.314 4 24s2.68-6 5.984-6h2.005M24 18h13.989A6.006 6.006 0 0 1 44 24c0 3.314-2.691 6-6.011 6h-1.923"/></g>`),
		g.Group(children),
	)
}