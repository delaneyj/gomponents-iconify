package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RatingTwentyOnePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m10 3V9m2.5 3h3M17 10.5v3"/><path d="M7 10.5a1.5 1.5 0 0 1 3 0c0 .443-.313.989-.612 1.393L7 15h3"/></g>`),
		g.Group(children),
	)
}