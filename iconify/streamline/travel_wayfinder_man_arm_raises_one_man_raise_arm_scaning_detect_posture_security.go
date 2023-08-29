package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelWayfinderManArmRaisesOneManRaiseArmScaningDetectPostureSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="2"/><path d="M2.5 6.5h9m-3 0v4h-3v-4m0 4v3m3-3v3"/></g>`),
		g.Group(children),
	)
}