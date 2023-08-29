package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ppt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M4 8H44"/><path fill="#2F88FF" fill-rule="evenodd" stroke="#000" d="M8 8H40V34H8V8Z" clip-rule="evenodd"/><path stroke="#fff" d="M22 16L27 21L22 26"/><path stroke="#000" d="M16 42L24 34L32 42"/></g>`),
		g.Group(children),
	)
}