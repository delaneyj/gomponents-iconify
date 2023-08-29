package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h17v1H3V4m0 4h13v1H3V8m8 13v-8.25L7.75 16L7 15.34l4.5-4.5l4.5 4.5l-.75.66L12 12.75V21h-1Z"/>`),
		g.Group(children),
	)
}