package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderManArmRaisesTwoManRaiseArmScaningDetectPostureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M2.5 5a7.5 7.5 0 0 0 9 0"/><path d="M8.5 6.33v4.17h-3V6.33m0 4.17v3m3-3v3"/></g>`),
		g.Group(children),
	)
}