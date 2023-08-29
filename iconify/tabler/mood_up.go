package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.984 12.536a9 9 0 1 0-8.463 8.449M19 22v-6m3 3l-3-3l-3 3m-7-9h.01M15 10h.01"/><path d="M9.5 15c.658.64 1.56 1 2.5 1s1.842-.36 2.5-1"/></g>`),
		g.Group(children),
	)
}