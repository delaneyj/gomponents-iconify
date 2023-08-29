package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandMessenger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 20l1.3-3.9A9 8 0 1 1 7.7 19L3 20"/><path d="m8 13l3-2l2 2l3-2"/></g>`),
		g.Group(children),
	)
}