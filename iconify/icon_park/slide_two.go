package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M4 8H44"/><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M8 8H40V34H8V8Z" clip-rule="evenodd"/><path stroke="#fff" d="M31 18L34 21L31 24"/><path stroke="#fff" d="M17 24L14 21L17 18"/><path stroke="#000" d="M16 42L24 34L32 42"/></g>`),
		g.Group(children),
	)
}