package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.5 13.5L7 12l1.5-1.5m7 0L17 12l-1.5 1.5"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m10-2.5L11 15"/></g>`),
		g.Group(children),
	)
}