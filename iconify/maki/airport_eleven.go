package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirportEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.5 6.4V6l4.5.5V5L6.5 3.2V1.5C6.5.5 6 0 5.5 0s-1 .5-1 1.5v1.7L0 5v1.4L4.5 6v3.3L3 10v1l2.5-.5L8 11v-1l-1.5-.8V6.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}