package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevatorPassengerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.41-1.41v-.01L2.39 1.73L1.11 3L3 4.9V19c0 1.1.9 2 2 2h14.1l1.74 1.73l1.27-1.27M11 14h-1v4H7v-4H6v-2.5c0-1 .71-1.79 1.65-1.96L11 12.89V14m2.46-3.74L6.2 3H19c1.1 0 2 .9 2 2v12.8l-3.69-3.69L18 13h-1.8l-2-2H18l-2.5-4l-2.04 3.26Z"/>`),
		g.Group(children),
	)
}