package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiFourTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.963 13.787a1.298 1.298 0 1 1-1.836 1.836a1.298 1.298 0 0 1 1.836-1.836Z"/>`),
		g.Group(children),
	)
}