package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandDribbble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M9 3.6c5 6 7 10.5 7.5 16.2"/><path d="M6.4 19c3.5-3.5 6-6.5 14.5-6.4M3.1 10.75c5 0 9.814-.38 15.314-5"/></g>`),
		g.Group(children),
	)
}