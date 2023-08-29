package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m216.13 106.72l-80-75.54l-.05-.05a12 12 0 0 0-16.2.05l-79.95 75.49a12 12 0 0 0-3.93 8.87V208a12 12 0 0 0 12 12h160a12 12 0 0 0 12-12v-92.46a12 12 0 0 0-3.87-8.82ZM212 208a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-92.46a4.09 4.09 0 0 1 1.36-3l79.94-75.49a4 4 0 0 1 5.36 0l80 75.55a4 4 0 0 1 1.31 3Z"/>`),
		g.Group(children),
	)
}