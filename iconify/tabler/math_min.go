package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MathMin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 17a2 2 0 1 1 0 4a2 2 0 0 1 0-4z"/><path d="M3 4c0 8.75 4 14 7 14.5m4 0c3-.5 7-5.75 7-14.5"/></g>`),
		g.Group(children),
	)
}