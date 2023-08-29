package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFlashOffFlashPowerConnectChargeOffElectricityLightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.5.5l13 13M5.19 5.19L8.5.5V5h3L8.81 8.81m-1.66 2.35L5.5 13.5V9h-3l1.41-2"/>`),
		g.Group(children),
	)
}