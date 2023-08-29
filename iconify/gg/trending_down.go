package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.851 8.106L.437 9.52l7.07 7.072l6.365-6.364l4.242 4.242l-1.742 1.743l6.692 1.793l-1.793-6.692l-1.742 1.742l-5.657-5.657l-6.364 6.364L1.85 8.106Z"/>`),
		g.Group(children),
	)
}