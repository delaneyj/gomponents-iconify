package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M20 8L24 4M24 4L28 8M24 4V16"/><path d="M20 40L24 44M24 44L28 40M24 44V32"/><path d="M40 20L44 24M44 24L40 28M44 24H32"/><path d="M8 20L4 24M4 24L8 28M4 24H16"/><circle cx="24" cy="24" r="2"/></g>`),
		g.Group(children),
	)
}