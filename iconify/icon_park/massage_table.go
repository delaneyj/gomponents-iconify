package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MassageTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M14 17C14 15.8954 14.8954 15 16 15H42C43.1046 15 44 15.8954 44 17V23H14V17Z"/><path stroke-linecap="round" d="M26 23L14 37"/><path stroke-linecap="round" d="M32 23L44 37"/><path stroke-linecap="round" d="M14 23L6 23"/><path stroke-linecap="round" d="M39 31L19 31"/><path stroke-linecap="round" d="M6 13V23"/><path stroke-linecap="round" d="M14 23V40"/><path stroke-linecap="round" d="M44 23V40"/><path stroke-linecap="round" d="M9 14L3 12"/></g>`),
		g.Group(children),
	)
}