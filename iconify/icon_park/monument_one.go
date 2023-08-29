package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonumentOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="20" height="6" x="14" y="38"/><path fill="#2F88FF" d="M18 38L20 9L24 4L28 9L30 38H18Z"/><path stroke-linecap="round" d="M4 44H44"/></g>`),
		g.Group(children),
	)
}