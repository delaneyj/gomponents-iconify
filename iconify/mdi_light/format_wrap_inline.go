package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapInline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4v1H3V4h17m0 12v1h-6v-1h6m0 4v1H3v-1h17M7.5 8l4.5 9H3l4.5-9m0 2.24L4.62 16h5.76L7.5 10.24Z"/>`),
		g.Group(children),
	)
}