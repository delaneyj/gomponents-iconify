package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSouthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9h-2v6.59L5.41 4L4 5.41L15.59 17H9v2h10V9z"/>`),
		g.Group(children),
	)
}