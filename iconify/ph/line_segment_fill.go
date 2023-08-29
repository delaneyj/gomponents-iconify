package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSegmentFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214.64 86.62a32.07 32.07 0 0 1-38.89 4.94l-84.19 84.19a32 32 0 1 1-50.2-6.37a32.06 32.06 0 0 1 38.89-4.94l84.19-84.19a32 32 0 1 1 50.2 6.37Z"/>`),
		g.Group(children),
	)
}