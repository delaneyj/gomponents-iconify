package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="16" r="3"/><path d="M21 13v6M3 19V9a4 4 0 0 1 4-4a4 4 0 0 1 4 4v10m-8-6h8"/></g>`),
		g.Group(children),
	)
}