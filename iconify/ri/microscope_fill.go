package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroscopeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.196 2.268l3.25 5.63a1 1 0 0 1-.366 1.365l-1.3.75l1.001 1.732l-1.732 1l-1-1.733l-1.299.751a1 1 0 0 1-1.366-.366L8.546 8.215a5.002 5.002 0 0 0-3.222 6.56A4.976 4.976 0 0 1 8 14c1.684 0 3.174.833 4.08 2.109l7.688-4.439l1 1.732l-7.878 4.549A5.022 5.022 0 0 1 12.9 20H21v2l-17 .001A4.978 4.978 0 0 1 3 19c0-1.007.298-1.945.81-2.73a7.003 7.003 0 0 1 3.717-9.82l-.393-.682a2 2 0 0 1 .732-2.732l2.598-1.5a2 2 0 0 1 2.732.732Z"/>`),
		g.Group(children),
	)
}