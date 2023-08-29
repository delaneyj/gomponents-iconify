package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 23.5A1.5 1.5 0 0 1 7.5 22h33a1.5 1.5 0 0 1 0 3h-33A1.5 1.5 0 0 1 6 23.5Z"/>`),
		g.Group(children),
	)
}