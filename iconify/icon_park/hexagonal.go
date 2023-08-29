package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M7 14L41 14"/><path d="M7 34L24 4"/><path d="M41 34L24 4"/><path d="M41 34L7 34"/><path d="M41 14L24 44"/><path d="M7 14L24 44"/><path d="M15.5 9L7 14V24V34L15.5 39L24 44L32.5001 39L41 34V24V14L32.5001 9L24 4L15.5 9Z"/></g>`),
		g.Group(children),
	)
}