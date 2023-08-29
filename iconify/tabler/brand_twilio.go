package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandTwilio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="M8 9a1 1 0 1 0 2 0a1 1 0 1 0-2 0m6 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0m0 6a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-6 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}