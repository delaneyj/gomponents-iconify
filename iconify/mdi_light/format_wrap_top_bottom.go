package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapTopBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.5 8l4.5 9H7l4.5-9m0 2.24L8.62 16h5.76l-2.88-5.76M3 4h17v1H3V4m17 16v1H3v-1h17Z"/>`),
		g.Group(children),
	)
}