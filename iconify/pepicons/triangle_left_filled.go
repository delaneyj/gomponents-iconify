package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m4.205 8.728l8-5A1.5 1.5 0 0 1 14.5 5v10a1.5 1.5 0 0 1-2.295 1.272l-8-5a1.5 1.5 0 0 1 0-2.544Z"/>`),
		g.Group(children),
	)
}