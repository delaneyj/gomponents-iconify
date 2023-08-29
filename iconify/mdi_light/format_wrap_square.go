package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.5 8l4.5 9H7l4.5-9m0 2.24L8.62 16h5.76l-2.88-5.76M3 16h2v1H3v-1m0-4h2v1H3v-1m0-4h2v1H3V8m15 0h2v1h-2V8m0 4h2v1h-2v-1m0 4h2v1h-2v-1m2 4v1H3v-1h17M3 4h17v1H3V4Z"/>`),
		g.Group(children),
	)
}