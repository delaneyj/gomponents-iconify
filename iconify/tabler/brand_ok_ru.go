package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandOkRu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 9a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M20 12c0 8 0 8-8 8s-8 0-8-8s0-8 8-8s8 0 8 8z"/><path d="M9.5 13c1.333.667 3.667.667 5 0m-5 4l2.5-3l2.5 3M12 13.5v.5"/></g>`),
		g.Group(children),
	)
}