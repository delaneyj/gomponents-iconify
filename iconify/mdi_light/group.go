package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 8v5h6V8H7M2 3h3v1h13V3h3v3h-1v13h1v3h-3v-1H5v1H2v-3h1V6H2V3m3 16v1h13v-1h1V6h-1V5H5v1H4v13h1M6 7h8v4h3v7H8v-4H6V7m8 7H9v3h7v-5h-2v2M3 4v1h1V4H3m16 0v1h1V4h-1m0 16v1h1v-1h-1M3 20v1h1v-1H3Z"/>`),
		g.Group(children),
	)
}