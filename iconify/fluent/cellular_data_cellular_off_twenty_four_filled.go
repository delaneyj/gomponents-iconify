package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellularDataCellularOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M3.28 2.22a.75.75 0 1 0-1.06 1.06L11 12.06V19a1 1 0 1 0 2 0v-4.94l2 2v2.945c0 .55.448.995 1 .995s1-.446 1-.995v-.944l3.72 3.72a.75.75 0 0 0 1.06-1.061L3.28 2.22z" fill="currentColor"/><path d="M19 15.818l2 2V6a1 1 0 0 0-2 0v9.82z" fill="currentColor"/><path d="M15 11.818l2 2V8.995A.998.998 0 0 0 16 8c-.552 0-1 .446-1 .995v2.823z" fill="currentColor"/><path d="M8 14c.552 0 1 .445 1 .994v4.012A.997.997 0 0 1 8 20c-.552 0-1-.445-1-.994v-4.012c0-.549.448-.994 1-.994z" fill="currentColor"/><path d="M4 17c.552 0 1 .44 1 .984v1.032A.992.992 0 0 1 4 20c-.552 0-1-.44-1-.984v-1.032A.992.992 0 0 1 4 17z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}