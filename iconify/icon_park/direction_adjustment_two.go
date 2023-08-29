package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustmentTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 10L24 4M24 4L30 10M24 4V14"/><path d="M18 38L24 44M24 44L30 38M24 44V34"/><path d="M38 18L44 24M44 24L38 30M44 24H34"/><path d="M10 18L4 24M4 24L10 30M4 24H14"/><circle cx="24" cy="24" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}