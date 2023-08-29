package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbMemoryStick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M8 22C8 20.8954 8.89543 20 10 20H38C39.1046 20 40 20.8954 40 22V42C40 43.1046 39.1046 44 38 44H10C8.89543 44 8 43.1046 8 42V22Z"/><path fill="#2F88FF" stroke="#000" d="M15 4H33V20H15V4Z"/><path stroke="#fff" d="M21 10V12"/><path stroke="#fff" d="M27 10V12"/></g>`),
		g.Group(children),
	)
}