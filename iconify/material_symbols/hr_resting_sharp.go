package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HrRestingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm2.425-9.6q-.725-.675-1.062-1.6t-.338-1.9q0-2.025 1.388-3.463T8.8 3q.9 0 1.725.325t1.475.95q.65-.625 1.475-.95T15.2 3q2.025 0 3.413 1.45T20 7.925q0 .975-.363 1.863T18.6 11.375l-6.6 6.4L5.425 11.4Z"/>`),
		g.Group(children),
	)
}