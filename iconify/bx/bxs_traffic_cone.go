package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsTrafficCone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M18.649 16H5.352l-1.06 3H2v2h20v-2h-2.292zM6.057 14h11.886l-1.412-4H7.469zM13 2h-2a1 1 0 0 0-.943.667L8.175 8h7.65l-1.882-5.333A1 1 0 0 0 13 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}