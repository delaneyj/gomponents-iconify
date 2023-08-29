package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SalesReport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M41 13.9997L24 4L7 13.9997V33.9998L24 44L41 33.9998V13.9997Z"/><path stroke="#fff" stroke-linecap="round" d="M24 22V30"/><path stroke="#fff" stroke-linecap="round" d="M32 18V30"/><path stroke="#fff" stroke-linecap="round" d="M16 26V30"/></g>`),
		g.Group(children),
	)
}