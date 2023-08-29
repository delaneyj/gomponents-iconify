package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataFiveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 17c.552 0 1 .44 1 .984v1.032A.992.992 0 0 1 4 20c-.552 0-1-.44-1-.984v-1.032A.992.992 0 0 1 4 17Z"/>`),
		g.Group(children),
	)
}