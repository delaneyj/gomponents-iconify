package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRollerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 104v50a16.07 16.07 0 0 1-11.6 15.38L136 198v34a8 8 0 0 1-16 0v-34a16.07 16.07 0 0 1 11.6-15.38L232 154v-50h-16v24a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16v-24H16a8 8 0 0 1 0-16h16V64a16 16 0 0 1 16-16h152a16 16 0 0 1 16 16v24h16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}