package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilkOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="26" height="22" x="6" y="22" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M14 38V28L19 34L24 28V38"/><path stroke="#000" d="M42 20L30 10"/><path stroke="#000" d="M20 6V12L30 10V4L20 6Z"/><path fill="#2F88FF" stroke="#000" d="M32 22L42 20V41L32 44V22Z"/><path stroke="#000" d="M19.4815 12L6 22H32L19.4815 12Z"/></g>`),
		g.Group(children),
	)
}