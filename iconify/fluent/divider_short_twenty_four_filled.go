package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividerShortTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 5v14a1 1 0 1 0 2 0V5a1 1 0 1 0-2 0Z"/>`),
		g.Group(children),
	)
}