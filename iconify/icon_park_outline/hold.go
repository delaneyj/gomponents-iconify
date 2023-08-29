package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m20 24l6 2s15-3 17-3s2 2 0 4s-9 8-15 8s-10-3-14-3H4"/><path d="M4 20c2-2 6-5 10-5s13.5 4 15 6s-3 5-3 5"/></g>`),
		g.Group(children),
	)
}