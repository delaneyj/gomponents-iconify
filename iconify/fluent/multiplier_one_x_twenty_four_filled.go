package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiplierOneXTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 7.902a1 1 0 0 0-1.854-.52l-.073.12c-.403.667-.842 1.394-2.087 2.14a1 1 0 0 0 1.029 1.716A8.53 8.53 0 0 0 8 10.672v5.329a1 1 0 1 0 2 0v-8.1Zm3.707 4.392a1 1 0 1 0-1.414 1.414l.793.793l-.793.793a1 1 0 0 0 1.414 1.414l.793-.793l.793.793a1 1 0 1 0 1.414-1.414l-.793-.793l.793-.793a1 1 0 0 0-1.414-1.414l-.793.793l-.793-.793Z"/>`),
		g.Group(children),
	)
}