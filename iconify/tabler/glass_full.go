package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 21h8m-4-6v6m5-18l1 7c0 3.012-2.686 5-6 5s-6-1.988-6-5l1-7h10z"/><path d="M6 10a5 5 0 0 1 6 0a5 5 0 0 0 6 0"/></g>`),
		g.Group(children),
	)
}