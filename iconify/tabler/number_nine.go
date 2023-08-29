package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 8a4 4 0 1 0-8 0v1a4 4 0 1 0 8 0"/><path d="M8 16a4 4 0 1 0 8 0V8"/></g>`),
		g.Group(children),
	)
}