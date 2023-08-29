package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundaboutLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21v-6.1q0-.725.475-1.275t1.2-.675q1.45-.225 2.388-1.35T20 9q0-1.65-1.175-2.825T16 5q-1.475 0-2.6.938t-1.35 2.387q-.125.725-.675 1.2T10.1 10H5.825l1.6 1.6L6 13L2 9l4-4l1.4 1.4L5.825 8h4.25q.35-2.2 2.038-3.6T16 3q2.5 0 4.25 1.75T22 9q0 2.2-1.4 3.888T17 14.924V21h-2Z"/>`),
		g.Group(children),
	)
}