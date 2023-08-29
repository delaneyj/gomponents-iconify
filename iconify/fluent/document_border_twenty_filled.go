package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentBorderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2A2.5 2.5 0 0 0 4 4.5v11A2.5 2.5 0 0 0 6.5 18h7a2.5 2.5 0 0 0 2.5-2.5v-11A2.5 2.5 0 0 0 13.5 2h-7ZM5 4.5A1.5 1.5 0 0 1 6.5 3h7A1.5 1.5 0 0 1 15 4.5v11a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 5 15.5v-11ZM6.75 4a.75.75 0 0 0-.75.75v10.5c0 .414.336.75.75.75h6.5a.75.75 0 0 0 .75-.75V4.75a.75.75 0 0 0-.75-.75h-6.5Z"/>`),
		g.Group(children),
	)
}