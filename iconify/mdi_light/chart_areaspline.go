package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAreaspline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h1v14l5.57-9.67l6.01 3.47l3.62-6.26l.86.5l-4.11 7.13L9.94 9.7L4 20h2.3l3.87-6.71l.5-.86l.86.5l5.15 2.97L20 10.14V21H3V4m14.04 13.26l-6.01-3.47L7.45 20H19v-6.12l-1.96 3.38Z"/>`),
		g.Group(children),
	)
}