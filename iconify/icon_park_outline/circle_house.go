package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 23H4s10.5-6 15-11s5.5-8 5.5-8s1 3 5.5 8s14 11 14 11ZM8 31h32v13H8zm5-8h22v8H13z"/>`),
		g.Group(children),
	)
}