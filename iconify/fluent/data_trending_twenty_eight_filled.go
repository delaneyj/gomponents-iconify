package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTrendingTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 0 1 1 1v17a2 2 0 0 0 2 2h17a1 1 0 1 1 0 2H7a4 4 0 0 1-4-4V4a1 1 0 0 1 1-1Zm19.002 6.412V12a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1h-5a1 1 0 1 0 0 2h2.584L16.5 13.086l-2.293-2.293a1 1 0 0 0-1.414 0l-5.5 5.5a1 1 0 1 0 1.414 1.414l4.793-4.793l2.293 2.293a1 1 0 0 0 1.414 0l5.795-5.795Z"/>`),
		g.Group(children),
	)
}