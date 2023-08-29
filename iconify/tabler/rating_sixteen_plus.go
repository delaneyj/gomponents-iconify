package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RatingSixteenPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0"/><path d="M10 13.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0-3 0M7 15V9m8.5 3h3M17 10.5v3"/><path d="M10 13.5v-3A1.5 1.5 0 0 1 11.5 9h1"/></g>`),
		g.Group(children),
	)
}