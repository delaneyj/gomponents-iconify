package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFlashTwoFlashPowerConnectChargeElectricityLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 .5v5h3.5l-5.5 8v-5H2.5L8 .5z"/>`),
		g.Group(children),
	)
}