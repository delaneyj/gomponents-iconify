package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerAltO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50.049 16.035a4.725 2.593 0 0 0-3.39.76L1.382 41.641a4.725 2.593 0 0 0 0 3.668l45.275 24.845a4.725 2.593 0 0 0 6.684 0L98.617 45.31a4.725 2.593 0 0 0 0-3.668L53.342 16.795a4.725 2.593 0 0 0-3.293-.76ZM50 22.295l38.596 21.18L50 64.655l-38.594-21.18ZM4.727 52.857l-3.344 1.834a4.725 2.593 0 0 0 0 3.668l45.275 24.846a4.725 2.593 0 0 0 6.684 0L98.617 58.36a4.725 2.593 0 0 0-.002-3.668l-3.342-1.834l-6.683 3.666l.004.002L50 77.705l-38.596-21.18l.004-.002z" color="currentColor"/>`),
		g.Group(children),
	)
}