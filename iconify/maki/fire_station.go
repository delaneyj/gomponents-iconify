package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireStation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 14c3.59 0 6.5-3 6.5-6.5c0-3-2.5-5.5-2.5-5.5l-1 3.5l-3-4.5l-3 4.5l-1-3.5S1 4.5 1 7.5C1 11 3.91 14 7.5 14Zm0-1.5A2.5 2.5 0 0 1 5 10c0-1.38 2.5-4.5 2.5-4.5S10 8.62 10 10a2.5 2.5 0 0 1-2.5 2.5Z"/>`),
		g.Group(children),
	)
}