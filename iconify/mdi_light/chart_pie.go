package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3h1a8 8 0 0 1 8 8v1h-9V3m1 8h7a7 7 0 0 0-7-7v7m-3 3h8a8 8 0 0 1-8 8a8 8 0 0 1-8-8a8 8 0 0 1 8-8v8m-1 1V7.07c-3.39.49-6 3.4-6 6.93a7 7 0 0 0 7 7c3.53 0 6.44-2.61 6.93-6H9Z"/>`),
		g.Group(children),
	)
}