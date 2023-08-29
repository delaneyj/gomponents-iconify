package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyTreeOneTreePlantPineTrianglePark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M11.57 7.84a2 2 0 0 1-.26 1.82a2 2 0 0 1-1.63.84H4.32a2 2 0 0 1-1.63-.84a2 2 0 0 1-.26-1.82l1.91-5.45a2.82 2.82 0 0 1 5.32 0Z"/><path d="M5.5 6.5L7 8v5.5M7 8l1.5-1.5"/></g>`),
		g.Group(children),
	)
}