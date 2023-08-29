package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M15.5 9L7 14V24V34L15.5 39L24 44L32.5001 39L41 34V24V14L32.5001 9L24 4L15.5 9Z"/><path d="M41 14L24 24"/><path d="M7 14L24 24"/><path d="M24 44V24"/><path d="M32 19L32 39"/><path d="M41 24L24 34"/><path d="M24 34L7 24"/><path d="M16 39L16 19"/><path d="M32 9L16 19"/><path d="M32 19L16 9"/></g>`),
		g.Group(children),
	)
}