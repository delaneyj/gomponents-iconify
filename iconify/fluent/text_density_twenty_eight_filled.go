package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDensityTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.25 2a.75.75 0 0 1 .75.75v22.5a.75.75 0 0 1-1.5 0V2.75a.75.75 0 0 1 .75-.75ZM12 5H4a.75.75 0 0 0 0 1.5h8V5Zm0 4H4a.75.75 0 0 0 0 1.5h8V9Zm0 4H4a.75.75 0 0 0 0 1.5h8V13Zm0 4H4a.75.75 0 0 0 0 1.5h8V17Zm0 4H4a.75.75 0 0 0 0 1.5h8V21Zm10.5 0h-6v-5h6a2.5 2.5 0 0 1 0 5Zm0-9h-6V7h6a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}