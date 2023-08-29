package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instruction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" stroke-linejoin="round" rx="2"/><path fill="#2F88FF" stroke-linejoin="round" d="M4 10C4 8.89543 4.89543 8 6 8H42C43.1046 8 44 8.89543 44 10V16H4V10Z"/><path d="M25 23L23 34"/><path stroke-linejoin="round" d="M31 23L37 28L31 34"/><path stroke-linejoin="round" d="M17 22.9999L11 27.9999L17 33.9999"/></g>`),
		g.Group(children),
	)
}