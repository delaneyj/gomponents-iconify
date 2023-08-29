package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleSwipeUpTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.207 2.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 1.414 1.414L16.5 5.414V17a1 1 0 1 0 2 0V5.414l1.293 1.293a1 1 0 1 0 1.414-1.414l-3-3ZM12.5 17a5.001 5.001 0 0 1 3-4.584v2.348a3 3 0 1 0 4 0v-2.348a5.001 5.001 0 0 1-2 9.584a5 5 0 0 1-5-5Zm-8-4.584A5.001 5.001 0 0 0 6.5 22a5 5 0 0 0 2-9.584v2.348a3 3 0 1 1-4 0v-2.348ZM7.207 2.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 1.414 1.414L5.5 5.414V17a1 1 0 1 0 2 0V5.414l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3Z"/>`),
		g.Group(children),
	)
}