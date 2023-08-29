package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 12.998a.75.75 0 0 0 0 1.5h22.5a.75.75 0 0 0 0-1.5H2.75ZM21.25 2A2.75 2.75 0 0 1 24 4.75v7.248H4V4.75A2.75 2.75 0 0 1 6.75 2h14.5ZM4 22.75v-7.252h20v7.252a2.75 2.75 0 0 1-2.75 2.75H6.75A2.75 2.75 0 0 1 4 22.75Z"/>`),
		g.Group(children),
	)
}