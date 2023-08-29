package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAlpineJs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 11.5L7.5 16h9l-9-9z"/><path d="m16.5 16l4.5-4.5L16.5 7L12 11.5"/></g>`),
		g.Group(children),
	)
}