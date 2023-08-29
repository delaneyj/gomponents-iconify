package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTextMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M18 36L24 30L30 36"/><path d="M23.9999 30.9998V43.9998"/><path d="M18 12L24 18L30 12"/><path d="M23.9999 17.0002V4.00022"/><path d="M6 24.0004H42"/></g>`),
		g.Group(children),
	)
}