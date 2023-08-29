package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm15.5 5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-2.5.5a2.5 2.5 0 1 1 1.026 2.02l-8.041 4.696a2.5 2.5 0 1 1-1.003-1.73l8.035-4.693A2.534 2.534 0 0 1 17 7.5Zm-9.067 6.75a.5.5 0 1 0-.866.5a.5.5 0 0 0 .866-.5Z"/>`),
		g.Group(children),
	)
}