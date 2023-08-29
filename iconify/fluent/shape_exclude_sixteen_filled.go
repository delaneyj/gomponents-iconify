package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeExcludeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.268A2 2 0 0 0 2 4v5a2 2 0 0 0 2 2h1v1a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-1V4a2 2 0 0 0-2-2H4a1.991 1.991 0 0 0-1 .268ZM11 5v4a2 2 0 0 1-2 2H5V7a2 2 0 0 1 2-2h4Z"/>`),
		g.Group(children),
	)
}