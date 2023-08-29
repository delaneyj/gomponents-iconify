package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineStacked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.45 15.18L22 6.81V21H2V3h2v16h2.57l4.39-7.56l6.49 3.74M22 3l-.03.45L17 11l-7-5l-4 6V3h16Z"/>`),
		g.Group(children),
	)
}