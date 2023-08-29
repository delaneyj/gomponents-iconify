package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligiousJewish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5.5L5.5 4H1l2.5 3.5L1 11h4.5l2 3.5l2-3.5H14l-2.5-3.5L14 4H9.5L7.5.5Z"/>`),
		g.Group(children),
	)
}