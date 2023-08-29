package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DismissTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22.293 4.293a1 1 0 1 1 1.414 1.414L15.414 14l8.293 8.293a1 1 0 0 1-1.414 1.414L14 15.414l-8.293 8.293a1 1 0 0 1-1.414-1.414L12.586 14L4.293 5.707a1 1 0 0 1 1.414-1.414L14 12.586l8.293-8.293Z"/>`),
		g.Group(children),
	)
}