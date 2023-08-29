package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLightingBrightnessTwoBrightAdjustBrightnessAdjustmentSunRaiseControlsHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5v1m0 11v1M1.5 7h-1m2.61 3.89l-.71.71m.71-8.49L2.4 2.4M7 3.5a3.5 3.5 0 0 0 0 7ZM12.5 7h1m-2.61 3.89l.71.71m-.71-8.49l.71-.71"/>`),
		g.Group(children),
	)
}