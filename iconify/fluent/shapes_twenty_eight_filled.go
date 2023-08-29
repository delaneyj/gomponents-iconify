package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapesTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 14v3.938A8.001 8.001 0 0 1 10 2a8 8 0 0 1 7.965 7.25H13.75A4.75 4.75 0 0 0 9 14Zm4.75-3.75A3.75 3.75 0 0 0 10 14v8.5a3.75 3.75 0 0 0 3.75 3.75h8.5A3.75 3.75 0 0 0 26 22.5V14a3.75 3.75 0 0 0-3.75-3.75h-8.5Z"/>`),
		g.Group(children),
	)
}