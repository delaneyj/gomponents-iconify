package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCompassTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 10a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0ZM4 4.5A2.5 2.5 0 0 1 6.5 2H18a2.5 2.5 0 0 1 2.5 2.5v14.25a.75.75 0 0 1-.75.75H5.5a1 1 0 0 0 1 1h13.25a.75.75 0 0 1 0 1.5H6.5A2.5 2.5 0 0 1 4 19.5v-15Zm9 1.25a.75.75 0 0 0-1.5 0v1.604a2.751 2.751 0 0 0-1.152 4.632l-1.294 3.236a.75.75 0 0 0 1.392.556l1.235-3.087a2.76 2.76 0 0 0 1.138 0l1.235 3.087a.75.75 0 0 0 1.392-.556l-1.294-3.236A2.751 2.751 0 0 0 13 7.354V5.75Z"/>`),
		g.Group(children),
	)
}