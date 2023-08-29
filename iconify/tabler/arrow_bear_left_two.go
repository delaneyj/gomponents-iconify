package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBearLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 3H4v5"/><path d="m4 3l7.536 7.536A5 5 0 0 1 13 14.07V21m7-16l-4.5 4.5"/></g>`),
		g.Group(children),
	)
}