package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatFloatLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v1H3V4m9 4h8v1h-8V8m0 4h8v1h-8v-1m-9 4h13v1H3v-1m17 4v1H3v-1h17M3 7h7v7H3V7m6 1H4v5h5V8Z"/>`),
		g.Group(children),
	)
}