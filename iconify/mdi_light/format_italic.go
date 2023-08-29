package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatItalic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17v-1h3l4-11h-3V4h7v1h-3l-4 11h3v1H6Z"/>`),
		g.Group(children),
	)
}