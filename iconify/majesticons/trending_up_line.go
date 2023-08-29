package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2.293 17.707a1 1 0 0 0 1.414 0L9 12.414l3.293 3.293a1 1 0 0 0 1.414 0L20 9.414V14a1 1 0 1 0 2 0V7a1 1 0 0 0-1-1h-7a1 1 0 1 0 0 2h4.586L13 13.586l-3.293-3.293a1 1 0 0 0-1.414 0l-6 6a1 1 0 0 0 0 1.414z"/></g>`),
		g.Group(children),
	)
}