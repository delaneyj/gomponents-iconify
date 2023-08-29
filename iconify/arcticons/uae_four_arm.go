package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UaeFourArm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.25 28l8.16 11.89l21.26-31.8M18.68 35.82l2.81 4.09L42.75 8.09M13.33 28l2.74 4"/>`),
		g.Group(children),
	)
}