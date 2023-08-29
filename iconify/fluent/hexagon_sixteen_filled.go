package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.4 2a1.5 1.5 0 0 0-1.3.75l-2.599 4.5a1.5 1.5 0 0 0 0 1.5l2.6 4.5a1.5 1.5 0 0 0 1.298.75h5.2a1.5 1.5 0 0 0 1.3-.75l2.599-4.5a1.5 1.5 0 0 0 0-1.5l-2.6-4.5A1.5 1.5 0 0 0 10.6 2H5.4Z"/>`),
		g.Group(children),
	)
}