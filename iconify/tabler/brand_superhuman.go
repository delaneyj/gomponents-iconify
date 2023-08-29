package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSuperhuman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 12l4 3l-8 7l-8-7l4-3"/><path d="M12 3L4 9l8 6l8-6zm0 12h8"/></g>`),
		g.Group(children),
	)
}