package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandVite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 4.5L16 3l-2 6.5l2-.5l-4 7v-5l-3 1z"/><path d="M15 6.5L22 5L12 22L2 5l7.741 1.5"/></g>`),
		g.Group(children),
	)
}