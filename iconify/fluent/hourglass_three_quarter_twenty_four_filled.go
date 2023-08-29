package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassThreeQuarterTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.97 7c.24.397.565.738.95 1L10 9.48c.42.278.77.65 1.02 1.089a2.805 2.805 0 0 1 0 2.862c-.25.438-.6.811-1.02 1.089L7.92 16a3.24 3.24 0 0 0-1.42 2.65v.6a.25.25 0 0 0 .25.25h10.5a.25.25 0 0 0 .25-.25v-.6A3.24 3.24 0 0 0 16.08 16L14 14.52a3.22 3.22 0 0 1-1.02-1.089a2.805 2.805 0 0 1 0-2.862c.25-.438.6-.811 1.02-1.089L16.08 8a3.24 3.24 0 0 0 .95-1H6.97Z"/>`),
		g.Group(children),
	)
}