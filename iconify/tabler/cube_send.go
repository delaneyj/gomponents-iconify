package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeSend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 12.5l-5-3l5-3l5 3V15l-5 3z"/><path d="M11 9.5V15l5 3m0-5.455l5-3.03M7 9H2m5 3H4m3 3H6"/></g>`),
		g.Group(children),
	)
}