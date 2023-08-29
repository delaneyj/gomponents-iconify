package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineMulti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm17.914 3.5L15.5 11.914l-4-4l-5 5L5.086 11.5L11.5 5.086l4 4l5-5L21.914 5.5Zm0 5L15.5 16.914l-4-4l-5 5L5.086 16.5l6.414-6.414l4 4l5-5l1.414 1.414Z"/>`),
		g.Group(children),
	)
}