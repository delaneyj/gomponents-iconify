package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v1H3V4m4 4h9v1H7V8m-4 4h17v1H3v-1m4 4h9v1H7v-1m-4 4h17v1H3v-1Z"/>`),
		g.Group(children),
	)
}