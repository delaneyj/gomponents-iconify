package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SectionSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.172 19A3 3 0 1 0 12 15m2.83-10A3 3 0 1 0 12 9"/><path d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/></g>`),
		g.Group(children),
	)
}