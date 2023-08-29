package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandPocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 4h14a2 2 0 0 1 2 2v6a9 9 0 0 1-18 0V6a2 2 0 0 1 2-2"/><path d="m8 11l4 4l4-4"/></g>`),
		g.Group(children),
	)
}