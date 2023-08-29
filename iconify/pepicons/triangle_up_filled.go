package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.272 5.205l5 8A1.5 1.5 0 0 1 15 15.5H5a1.5 1.5 0 0 1-1.272-2.295l5-8a1.5 1.5 0 0 1 2.544 0Z"/>`),
		g.Group(children),
	)
}