package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareMultipleFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.497 10h-2.503a6.25 6.25 0 0 1 6.245-6h14.5a9.25 9.25 0 0 1 9.25 9.25v14.5A6.25 6.25 0 0 1 38 33.995V31.49a3.75 3.75 0 0 0 3.49-3.741v-14.5a6.75 6.75 0 0 0-6.75-6.75h-14.5a3.75 3.75 0 0 0-3.743 3.5Zm-6.247 2A6.25 6.25 0 0 0 4 18.25v19.5A6.25 6.25 0 0 0 10.25 44h19.5A6.25 6.25 0 0 0 36 37.75v-19.5A6.25 6.25 0 0 0 29.75 12h-19.5ZM6.5 18.25a3.75 3.75 0 0 1 3.75-3.75h19.5a3.75 3.75 0 0 1 3.75 3.75v19.5a3.75 3.75 0 0 1-3.75 3.75h-19.5a3.75 3.75 0 0 1-3.75-3.75v-19.5Z"/>`),
		g.Group(children),
	)
}