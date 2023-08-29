package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 96H188l-43-43h324v43zM41 3l42 42l372 373l-27 27l-50-50H0v-64h43V96q0-15 10-27L14 30zm44 99v229h229zm406 37q8 0 14.5 6t6.5 15v213q0 9-6.5 15.5T491 395h-4l-64-64h46V181h-85v111l-43-43v-89q0-9 6.5-15t15.5-6h128z"/>`),
		g.Group(children),
	)
}