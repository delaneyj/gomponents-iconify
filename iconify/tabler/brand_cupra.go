package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCupra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.5 10L2 6l15.298 6.909a.2.2 0 0 1 .09.283L14 19"/><path d="m10 19l-3.388-5.808a.2.2 0 0 1 .09-.283L22 6l-2.5 4"/></g>`),
		g.Group(children),
	)
}