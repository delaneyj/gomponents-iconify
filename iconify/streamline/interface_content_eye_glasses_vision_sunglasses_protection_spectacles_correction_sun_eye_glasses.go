package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceContentEyeGlassesVisionSunglassesProtectionSpectaclesCorrectionSunEyeGlasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 9h-5v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1Zm8 0h-5v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1Zm-8 0h3m5 0V3a2 2 0 0 0-2-2h-1M.5 9V3a2 2 0 0 1 2-2h1"/>`),
		g.Group(children),
	)
}