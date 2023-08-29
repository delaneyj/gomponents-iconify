package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlashForwardTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.316 2.051a1 1 0 0 1 .633 1.265l-6 18a1 1 0 1 1-1.897-.632l6-18a1 1 0 0 1 1.264-.633Z"/>`),
		g.Group(children),
	)
}