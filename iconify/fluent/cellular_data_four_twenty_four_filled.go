package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataFourTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14c.552 0 1 .445 1 .994v4.012c0 .549-.448.994-1 .994s-1-.445-1-.994v-4.012A.997.997 0 0 1 8 14Zm-4 3c.552 0 1 .44 1 .984v1.032A.992.992 0 0 1 4 20c-.552 0-1-.44-1-.984v-1.032A.992.992 0 0 1 4 17Z"/>`),
		g.Group(children),
	)
}