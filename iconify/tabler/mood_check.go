package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.925 13.163A8.998 8.998 0 0 0 12 3a9 9 0 0 0 0 18M9 10h.01M15 10h.01"/><path d="M9.5 15c.658.64 1.56 1 2.5 1s1.842-.36 2.5-1m.5 4l2 2l4-4"/></g>`),
		g.Group(children),
	)
}