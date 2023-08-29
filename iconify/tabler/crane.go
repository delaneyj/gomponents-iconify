package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 21h6m-3 0V3L3 9h18M9 3l10 6"/><path d="M17 9v4a2 2 0 1 1-2 2"/></g>`),
		g.Group(children),
	)
}