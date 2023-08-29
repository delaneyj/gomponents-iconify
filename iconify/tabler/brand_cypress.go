package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCypress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.48 17.007A9 9 0 1 0 12 21a2.08 2.08 0 0 0 1.974-1.423L17.5 9m-4 0l2 6"/><path d="M10.764 9.411a3 3 0 1 0-.023 5.19"/></g>`),
		g.Group(children),
	)
}