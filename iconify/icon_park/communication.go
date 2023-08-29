package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Communication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M33 38H22V30H36V22H44V38H39L36 41L33 38Z"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M4 6H36V30H17L13 34L9 30H4V6Z"/><path stroke="#fff" d="M19 18H20"/><path stroke="#fff" d="M26 18H27"/><path stroke="#fff" d="M12 18H13"/></g>`),
		g.Group(children),
	)
}