package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 17.707a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0L13 13.586l3.086-3.086l-2.793-2.793A1 1 0 0 1 14 6h7a1 1 0 0 1 1 1v7a1 1 0 0 1-1.707.707L17.5 11.914l-3.793 3.793a1 1 0 0 1-1.414 0L9 12.414l-5.293 5.293z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}