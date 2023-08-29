package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3 4h18a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1h1l2 3v1H0v-1l2-3h1a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m1 2v9h16V6H4z" fill="currentColor"/>`),
		g.Group(children),
	)
}