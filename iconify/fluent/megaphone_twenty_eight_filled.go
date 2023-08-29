package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M26 7.353a2.75 2.75 0 0 0-3.458-2.657L4.045 9.629a2.75 2.75 0 0 0-2.041 2.657v3.427a2.75 2.75 0 0 0 2.041 2.657L7 19.158v.342a4.5 4.5 0 0 0 8.56 1.942l6.982 1.862A2.75 2.75 0 0 0 26 20.647V7.352ZM8.5 19.56l5.572 1.486A3 3 0 0 1 8.5 19.559Z"/>`),
		g.Group(children),
	)
}