package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoLineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M6 7H42"/><path d="M6 41H42"/><path stroke-linejoin="round" d="M24 13L14 35"/><path stroke-linejoin="round" d="M18 28L30 28"/><path stroke-linejoin="round" d="M24 13L34 35"/></g>`),
		g.Group(children),
	)
}