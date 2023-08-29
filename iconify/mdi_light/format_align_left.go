package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-1h17v1H3m0-4v-1h11v1H3m0-4v-1h17v1H3m0-4V8h11v1H3m0-4V4h17v1H3Z"/>`),
		g.Group(children),
	)
}