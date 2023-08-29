package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 993q0 13-9.5 22.5T864 1025H160q-13 0-22.5-9.5T128 993q0-32-13-62.5t-32-58t-38-58T13 739T0 641V417q0-13 9.5-22.5T32 385h594L836 17q7-12 20-15.5t24 3T894.5 24T892 49L700 385h22l247-247q9-10 22.5-10t23 9.5t9.5 23t-10 22.5L812 385h180q13 0 22.5 9.5t9.5 22.5v224q0 53-13 98.5T979 816t-38 58.5t-32 57.5t-13 61z"/>`),
		g.Group(children),
	)
}