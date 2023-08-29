package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 7l4-4l4 4"/><path d="M14 3v4.394A6.737 6.737 0 0 1 11 13a6.737 6.737 0 0 0-3 5.606V21"/></g>`),
		g.Group(children),
	)
}