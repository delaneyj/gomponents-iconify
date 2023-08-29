package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.5 10l-3 5h4l-3 5.5M8 8.036a3.484 3.484 0 0 1 1.975.99M7.5 15a3.5 3.5 0 1 1 .87-6.891a5.002 5.002 0 0 1 9.614 1.49A2.751 2.751 0 0 1 17.25 15"/>`),
		g.Group(children),
	)
}