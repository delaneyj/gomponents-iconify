package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbTypeC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M15 15H33C48 15 48 33 33 33H15C-1.23978e-05 33 -2.86102e-06 15 15 15Z"/><path stroke="#fff" d="M21 27V21"/><path stroke="#fff" d="M27 27V21"/><path stroke="#fff" d="M33 27V21"/><path stroke="#fff" d="M15 27V21"/><path stroke="#fff" d="M36 24L12 24"/></g>`),
		g.Group(children),
	)
}