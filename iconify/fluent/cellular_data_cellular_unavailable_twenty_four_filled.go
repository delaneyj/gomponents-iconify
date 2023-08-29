package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataCellularUnavailableTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M20 19a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM16 9c.552 0 1 .446 1 .995v10.01c0 .55-.448.995-1 .995s-1-.446-1-.995V9.995c0-.55.448-.995 1-.995zm-4 3a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-7a1 1 0 0 1 1-1zm-4 3c.552 0 1 .445 1 .994v4.012c0 .549-.448.994-1 .994s-1-.445-1-.994v-4.012c0-.549.448-.994 1-.994zm-4 3c.552 0 1 .44 1 .984v1.032A.992.992 0 0 1 4 21c-.552 0-1-.44-1-.984v-1.032A.992.992 0 0 1 4 18zM20 6a1 1 0 0 1 1 1v9.996a1 1 0 0 1-2 0V7A1 1 0 0 1 20 6z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}