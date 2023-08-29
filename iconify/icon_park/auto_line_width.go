package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoLineWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M5 40L5 8"/><path stroke-linejoin="round" d="M24 13L14 35"/><path stroke-linejoin="round" d="M18 28L30 28"/><path stroke-linejoin="round" d="M24 13L34 35"/><path d="M43 40L43 8"/></g>`),
		g.Group(children),
	)
}