package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13v-1h17v1H3m8-9h1v4.25L14.25 6l.75.66l-3.5 3.5L8 6.66L8.75 6L11 8.25V4m0 17v-4.25L8.75 19L8 18.34l3.5-3.5l3.5 3.5l-.75.66L12 16.75V21h-1Z"/>`),
		g.Group(children),
	)
}