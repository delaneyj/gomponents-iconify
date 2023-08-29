package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatFloatRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4v1H3V4h17m-9 4v1H3V8h8m-8 4h5v1H3v-1m0 4h13v1H3v-1m0 4h17v1H3v-1M20 7v7h-7V7h7m-1 1h-5v5h5V8Z"/>`),
		g.Group(children),
	)
}