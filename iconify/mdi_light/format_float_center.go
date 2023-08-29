package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatFloatCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v1H3V4m0 12h17v1H3v-1m13 4v1H3v-1h13M8 7h7v7H8V7m6 1H9v5h5V8Z"/>`),
		g.Group(children),
	)
}