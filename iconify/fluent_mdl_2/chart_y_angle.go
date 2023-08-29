package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartYAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024h-293l18 19q137 143 206 315t69 370v64H37L1683 147l90 90L347 1664h1571q-6-75-25-151t-51-148t-76-135t-102-115v293h-128V896h512v128z"/>`),
		g.Group(children),
	)
}