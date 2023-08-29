package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustmentThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 10L24 4M24 4L30 10M24 4V24M24 24V44M24 24H44M24 24H4M18 38L24 44M24 44L30 38M38 18L44 24M44 24L38 30M10 18L4 24M4 24L10 30"/>`),
		g.Group(children),
	)
}