package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.553 15.536a4.999 4.999 0 0 1-7.902-1.098h2.381l.696-1.876H10.05a5.047 5.047 0 0 1 0-1.125h4.287l.696-1.874h-4.38a4.998 4.998 0 0 1 7.902-1.099l1.414-1.414A7.003 7.003 0 0 0 8.454 9.562H6.032l-.696 1.875H8.04a7.095 7.095 0 0 0 0 1.126H4.728l-.696 1.874h4.422a7.003 7.003 0 0 0 11.514 2.513l-1.415-1.414Z"/>`),
		g.Group(children),
	)
}