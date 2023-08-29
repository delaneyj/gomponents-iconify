package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 4.27L2.28 3L20 20.72L18.73 22l-2-2H0v-2h4a2 2 0 0 1-2-2V6c0-.22.04-.43.1-.63L1 4.27M4 16h8.73L4 7.27V16m16 0V6H7.82l-2-2H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h4v2h-2.18l-4-4H20Z"/>`),
		g.Group(children),
	)
}