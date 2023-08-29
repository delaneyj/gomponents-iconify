package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HangerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 4L24 44"/><path d="M32 4L24 10"/><path d="M16 9L24 15"/><path d="M16 20L24 26"/><path d="M32 15L24 21"/><path d="M30 44H18"/></g>`),
		g.Group(children),
	)
}