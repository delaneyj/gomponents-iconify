package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-1h13v1H3m0-4v-1h17v1H3m8-13h1v8.25L15.25 9l.75.66l-4.5 4.5L7 9.66L7.75 9L11 12.25V4Z"/>`),
		g.Group(children),
	)
}