package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<defs><path id="uiwPieChart0" d="M1.37 18.615h17.945c.379 0 .685.31.685.693a.689.689 0 0 1-.685.692H.685A.689.689 0 0 1 0 19.308V.692C0 .31.306 0 .685 0c.378 0 .684.31.684.692v17.923ZM2.836 17.4l2.753-5.57l2.754 1.392l3.442-4.179l2.754 2.09l4.131-6.268V17.4H2.837Z"/></defs><use fill="currentColor" href="#uiwPieChart0"/>`),
		g.Group(children),
	)
}