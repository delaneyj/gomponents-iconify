package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatFloatNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v1H3V4m9 9v-1h8v1h-8M3 7h7v7H3V7m1 1v5h5V8H4m-1 8h13v1H3v-1m0 4h17v1H3v-1Z"/>`),
		g.Group(children),
	)
}