package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 12l7 7m17 17l-7-7m7-17l-7 7M12 36l7-7m0-10h10v10H19zm17 0a7 7 0 1 0-7-7m7 17a7 7 0 1 1-7 7m-17-7a7 7 0 1 0 7 7m-7-17a7 7 0 1 1 7-7"/>`),
		g.Group(children),
	)
}