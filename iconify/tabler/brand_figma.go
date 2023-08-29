package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandFigma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M6 6a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v0a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3z"/><path d="M9 9a3 3 0 0 0 0 6h3m-3 0a3 3 0 1 0 3 3V3"/></g>`),
		g.Group(children),
	)
}