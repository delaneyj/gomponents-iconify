package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 3A1.5 1.5 0 0 0 5 4.5V28a1 1 0 1 0 2 0v-7h21a1 1 0 0 0 .8-1.6L23.25 12l5.55-7.4A1 1 0 0 0 28 3H6.5Z"/>`),
		g.Group(children),
	)
}