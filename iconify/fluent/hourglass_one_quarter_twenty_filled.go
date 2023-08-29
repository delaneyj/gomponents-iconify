package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassOneQuarterTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.68 14A3.41 3.41 0 0 0 6 16h8a3.49 3.49 0 0 0-.31-1.4a3.378 3.378 0 0 0-.367-.6H6.68Z"/>`),
		g.Group(children),
	)
}