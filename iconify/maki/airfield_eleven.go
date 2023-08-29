package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirfieldEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5 .5H3.5C3 .5 3 0 3.5 0h4c.5 0 .5.5 0 .5H6s.5.5.5 1.5v1H11v1.5l-4.5 2L6 10l1.5.5v.5h-4v-.5L5 10l-.5-3.5l-4.5-2V3h4.5V2C4.5 1 5 .5 5 .5z" fill="currentColor"/>`),
		g.Group(children),
	)
}