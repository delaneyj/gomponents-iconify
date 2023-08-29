package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CupOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 4.27L2.28 3L21 21.72L19.73 23l-1.46-1.46c-.34.29-.77.46-1.27.46H7a2.02 2.02 0 0 1-2-1.77L3.53 6.8L1 4.27M18.32 8l.45-4H5.82l-2-2H21l-1.71 15.47L9.82 8h8.5Z"/>`),
		g.Group(children),
	)
}