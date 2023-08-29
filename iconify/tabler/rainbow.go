package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 17c0-5.523-4.477-10-10-10S2 11.477 2 17"/><path d="M18 17a6 6 0 1 0-12 0"/><path d="M14 17a2 2 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}