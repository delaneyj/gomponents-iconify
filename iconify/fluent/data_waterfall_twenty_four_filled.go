package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataWaterfallTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 3a.75.75 0 0 0 0 1.5H4v5.75a2.25 2.25 0 0 0 2.25 2.25h6.25v6.25A2.25 2.25 0 0 0 14.75 21h6.5a.75.75 0 0 0 0-1.5H20v-6.25A2.25 2.25 0 0 0 17.75 11H11.5V5.25A2.25 2.25 0 0 0 9.25 3h-6.5Z"/>`),
		g.Group(children),
	)
}