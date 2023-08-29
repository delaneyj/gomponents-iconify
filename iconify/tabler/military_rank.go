package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilitaryRank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 7v13H7V7l5-3z"/><path d="m10 13l2-1l2 1m-4 4l2-1l2 1m-4-8l2-1l2 1"/></g>`),
		g.Group(children),
	)
}