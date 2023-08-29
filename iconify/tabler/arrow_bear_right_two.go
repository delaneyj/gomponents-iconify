package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBearRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 3h5v5"/><path d="m20 3l-7.536 7.536A5 5 0 0 0 11 14.07V21M4 5l4.5 4.5"/></g>`),
		g.Group(children),
	)
}