package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCoreos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0z"/><path d="M12 3c-3.263 3.212-3 7.654-3 12c4.59.244 8.814-.282 12-3"/><path d="M9.5 9a4.494 4.494 0 0 1 5.5 5.5"/></g>`),
		g.Group(children),
	)
}