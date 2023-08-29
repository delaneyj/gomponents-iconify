package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandLoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.464 6.518a6 6 0 1 0-3.023 7.965"/><path d="M17.482 17.464a6 6 0 1 0-7.965-3.023"/><path d="M6.54 17.482a6 6 0 1 0 3.024-7.965"/><path d="M6.518 6.54a6 6 0 1 0 7.965 3.024"/></g>`),
		g.Group(children),
	)
}