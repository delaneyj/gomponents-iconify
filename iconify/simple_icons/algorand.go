package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Algorand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.874 0h3.673l1.61 5.963h3.789l-2.588 4.5l3.624 13.533h-3.757l-2.44-9.077l-5.247 9.079H8.345l8.107-14.051l-1.304-4.878L4.215 24H.018Z"/>`),
		g.Group(children),
	)
}