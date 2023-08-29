package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3h12l4 6l-10 13L2 9Z"/><path d="M11 3L8 9l4 13l4-13l-3-6M2 9h20"/></g>`),
		g.Group(children),
	)
}