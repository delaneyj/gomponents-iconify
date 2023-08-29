package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 4.5a1 1 0 0 1 2 0v6.001c-.003 3.463 1.32 4.999 4.247 4.999c2.928 0 4.253-1.537 4.253-5v-6a1 1 0 1 1 2 0v6c0 4.54-2.18 7-6.253 7S5.996 15.039 6 10.5v-6ZM7 21a1 1 0 1 1 0-2h10.5a1 1 0 1 1 0 2H7Z"/>`),
		g.Group(children),
	)
}