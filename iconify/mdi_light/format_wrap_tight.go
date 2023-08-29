package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatWrapTight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.5 8l4.5 9H7l4.5-9m0 2.24L8.62 16h5.76l-2.88-5.76M3 4h17v1H3V4m11 4h6v1h-6V8M3 8h6v1H3V8m0 4h4v1H3v-1m0 4h2v1H3v-1m15 0h2v1h-2v-1m-2-4h4v1h-4v-1M3 20h17v1H3v-1Z"/>`),
		g.Group(children),
	)
}