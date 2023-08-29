package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBadgeDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M16.375 6.22L12 9.718l-4.375-3.5A1 1 0 0 0 6 7v6a1 1 0 0 0 .375.78l5 4a1 1 0 0 0 1.25 0l5-4A1 1 0 0 0 18 13V7a1 1 0 0 0-1.625-.78z"/></g>`),
		g.Group(children),
	)
}