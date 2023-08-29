package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMultipleFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.495 10h-2.503a6.25 6.25 0 0 1 6.245-6h14.5a9.25 9.25 0 0 1 9.25 9.25v14.5a6.25 6.25 0 0 1-5.989 6.245V31.49a3.75 3.75 0 0 0 3.49-3.741v-14.5a6.75 6.75 0 0 0-6.75-6.75h-14.5a3.75 3.75 0 0 0-3.742 3.5ZM4 18.25A6.25 6.25 0 0 1 10.25 12h19.5A6.25 6.25 0 0 1 36 18.25v19.5A6.25 6.25 0 0 1 29.75 44h-19.5A6.25 6.25 0 0 1 4 37.75v-19.5Z"/>`),
		g.Group(children),
	)
}