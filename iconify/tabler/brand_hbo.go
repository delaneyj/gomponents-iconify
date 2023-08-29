package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandHbo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 16V8m4 0v8m-4-4h4m3 4h2a2 2 0 1 0 0-4H9h2a2 2 0 1 0 0-4H9v8zm10-8a4 4 0 1 1 0 8a4 4 0 0 1 0-8z"/><path d="M18 12a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/></g>`),
		g.Group(children),
	)
}