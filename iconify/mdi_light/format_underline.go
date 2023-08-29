package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatUnderline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 11.5a5.5 5.5 0 0 1-5.5 5.5A5.5 5.5 0 0 1 6 11.5V4h1v7.5C7 14 9 16 11.5 16s4.5-2 4.5-4.5V4h1v7.5M5 21v-1h13v1H5Z"/>`),
		g.Group(children),
	)
}