package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDimensionalCodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M18 6H6V18H18V6Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M18 30H6V42H18V30Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M42 6H30V18H42V6Z"/><path stroke-linecap="round" d="M24 6V18"/><path stroke-linecap="round" d="M42 24H6"/><path stroke-linecap="round" d="M34 30V42"/><path stroke-linecap="round" d="M42 30V42"/><path stroke-linecap="round" d="M26 30V42"/></g>`),
		g.Group(children),
	)
}