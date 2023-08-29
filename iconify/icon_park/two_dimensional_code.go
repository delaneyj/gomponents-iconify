package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDimensionalCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M20 6H6V20H20V6Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M20 28H6V42H20V28Z"/><path fill="#2F88FF" stroke-linejoin="round" d="M42 6H28V20H42V6Z"/><path stroke-linecap="round" d="M29 28V42"/><path stroke-linecap="round" d="M41 28V42"/></g>`),
		g.Group(children),
	)
}