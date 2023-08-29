package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryCharge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 14H6C4.89543 14 4 14.8954 4 16V32C4 33.1046 4.89543 34 6 34H13"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 34H38C39.1046 34 40 33.1046 40 32V16C40 14.8954 39.1046 14 38 14H31"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22.002 14L17 24.0012H27.004L22 34"/><path fill="#000" d="M42 20H44C45.1046 20 46 20.8954 46 22V26C46 27.1046 45.1046 28 44 28H42V20Z"/></g>`),
		g.Group(children),
	)
}