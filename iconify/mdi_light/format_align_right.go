package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4v1H3V4h17m0 4v1H9V8h11m0 4v1H3v-1h17m0 4v1H9v-1h11m0 4v1H3v-1h17Z"/>`),
		g.Group(children),
	)
}