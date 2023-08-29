package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 6.293a1 1 0 0 0-1.414 1.414l6 6a1 1 0 0 0 1.414 0L13 10.414l3.086 3.086l-2.793 2.793A1 1 0 0 0 14 18h7a1 1 0 0 0 1-1v-7a1 1 0 0 0-1.707-.707L17.5 12.086l-3.793-3.793a1 1 0 0 0-1.414 0L9 11.586L3.707 6.293z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}