package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m10.44 2l.56.413v11.194l-.54.393L5 8.373v-.827L10.44 2z"/>`),
		g.Group(children),
	)
}