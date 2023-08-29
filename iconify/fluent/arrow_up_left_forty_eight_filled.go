package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeftFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.5 6a1.5 1.5 0 0 1 0 3H11.121l30.44 30.44a1.5 1.5 0 0 1-2.122 2.12L9 11.122V25.5a1.5 1.5 0 0 1-3 0v-18A1.5 1.5 0 0 1 7.5 6h18Z"/>`),
		g.Group(children),
	)
}