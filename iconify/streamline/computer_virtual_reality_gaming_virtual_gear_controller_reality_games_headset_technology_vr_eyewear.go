package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerVirtualRealityGamingVirtualGearControllerRealityGamesHeadsetTechnologyVrEyewear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 6.5H.5v5a1 1 0 0 0 1 1h3a1 1 0 0 0 .78-.38l1.31-1.63a.5.5 0 0 1 .78 0l1.33 1.63a1 1 0 0 0 .78.38h3a1 1 0 0 0 1-1Zm0 0L10.79 2a1 1 0 0 0-.86-.49H4.07a1 1 0 0 0-.86.49L.5 6.5"/>`),
		g.Group(children),
	)
}