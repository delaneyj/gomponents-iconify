package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFourTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.94 18.561a1.5 1.5 0 1 0 2.121-2.122a1.5 1.5 0 0 0-2.122 2.122Z"/>`),
		g.Group(children),
	)
}