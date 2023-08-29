package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookArrowClockwiseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.25 5a.75.75 0 0 0 .75-.75V1.5a.75.75 0 0 0-1.5 0v.626A6 6 0 1 0 23 7a.75.75 0 0 0-1.5 0a4.5 4.5 0 1 1-1.67-3.5H18.5a.75.75 0 0 0 0 1.5h2.75ZM6.5 2h5.601a7.028 7.028 0 0 0-1.165 1.5H6.5a1 1 0 0 0-1 1V18H19v-4.29a6.957 6.957 0 0 0 1.5-.647v5.687a.75.75 0 0 1-.75.75H5.5a1 1 0 0 0 1 1h13.25a.75.75 0 0 1 0 1.5H6.5A2.5 2.5 0 0 1 4 19.5v-15A2.5 2.5 0 0 1 6.5 2Z"/>`),
		g.Group(children),
	)
}