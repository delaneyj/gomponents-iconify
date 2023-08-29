package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyRainbowArchRainColorfulRainbowCurveHalfCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M.5 10.25a6.5 6.5 0 0 1 13 0"/><path d="M3 10.25a4 4 0 0 1 8 0"/><path d="M5.5 10.25a1.5 1.5 0 0 1 3 0"/></g>`),
		g.Group(children),
	)
}