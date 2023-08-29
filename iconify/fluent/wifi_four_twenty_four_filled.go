package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFourTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.96 18.567a1.501 1.501 0 1 0 2.122-2.122a1.501 1.501 0 0 0-2.123 2.122Z"/>`),
		g.Group(children),
	)
}