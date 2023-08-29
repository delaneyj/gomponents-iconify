package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="10" cy="14" r="8" opacity=".5"/><path d="M17 1.25a.75.75 0 0 0 0 1.5h3.19l-5.088 5.088c.385.32.74.674 1.06 1.06l5.088-5.087V7a.75.75 0 0 0 1.5 0V2.25a1 1 0 0 0-1-1H17Z"/></g>`),
		g.Group(children),
	)
}