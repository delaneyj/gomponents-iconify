package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.25 21.25a.75.75 0 0 0 1.5 0V20h2.5a2.25 2.25 0 0 0 2.25-2.25v-2.5A2.25 2.25 0 0 0 15.25 13h-2.5v-2h4.5a2.25 2.25 0 0 0 2.25-2.25v-2.5A2.25 2.25 0 0 0 17.25 4h-4.5V2.75a.75.75 0 0 0-1.5 0V4h-4.5A2.25 2.25 0 0 0 4.5 6.25v2.5A2.25 2.25 0 0 0 6.75 11h4.5v2h-2.5a2.25 2.25 0 0 0-2.25 2.25v2.5A2.25 2.25 0 0 0 8.75 20h2.5v1.25Z"/>`),
		g.Group(children),
	)
}