package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 5l7.6 5.5l-2.9 8.9H7.3l-2.9-8.9L12 5m0-2.5L2 9.8l3.8 11.7h12.3L22 9.8L12 2.5Z"/>`),
		g.Group(children),
	)
}