package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideTextMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.5 4.5A2.5 2.5 0 0 1 4 2h6.5A2.5 2.5 0 0 1 13 4.5v5a2.5 2.5 0 0 1-2.5 2.5H4a2.5 2.5 0 0 1-2.5-2.5v-5Zm3-.5a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1h-3Zm0 2a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1h-5ZM4 8.5a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 0-1h-3a.5.5 0 0 0-.5.5Zm6.5 4.5A3.5 3.5 0 0 0 14 9.5V3.671c.625.773 1 1.757 1 2.829v3a4.5 4.5 0 0 1-4.5 4.5h-4a4.481 4.481 0 0 1-2.829-1H10.5Z"/>`),
		g.Group(children),
	)
}