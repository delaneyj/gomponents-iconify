package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.5 6a1.5 1.5 0 0 0 0 3h14.379L6.439 39.44a1.5 1.5 0 1 0 2.122 2.12L39 11.122V25.5a1.5 1.5 0 0 0 3 0v-18A1.5 1.5 0 0 0 40.5 6h-18Z"/>`),
		g.Group(children),
	)
}