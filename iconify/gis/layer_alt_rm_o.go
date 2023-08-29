package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerAltRmO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M68.5 16.386v10.36h-37v-10.36zM24.045 29.207L1.383 41.641c-1.847 1.012-1.847 2.655 0 3.668l45.275 24.845c1.846 1.013 4.838 1.013 6.684 0L98.617 45.31c1.847-1.013 1.847-2.656 0-3.668L75.99 29.225l-.013.048l-4.588 4.76l17.207 9.442L50 64.655l-38.594-21.18l17.13-9.4ZM4.727 52.857l-3.344 1.834c-1.847 1.013-1.847 2.656 0 3.668l45.275 24.846c1.846 1.013 4.838 1.013 6.684 0L98.617 58.36c1.846-1.013 1.845-2.655-.002-3.668l-3.342-1.834l-6.683 3.666l.004.002L50 77.705l-38.596-21.18l.004-.002z" color="currentColor"/>`),
		g.Group(children),
	)
}