package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 4V44"/><path d="M16 9L16 37"/><path d="M24 9L24 37"/><path d="M32 9L32 37"/><path d="M42 44L6 44"/><path d="M42 37L6 37"/><path d="M39 4V44"/><path d="M42 9L6 9"/></g>`),
		g.Group(children),
	)
}