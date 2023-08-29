package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h3v1h8V3h3v3h-1v4h3V9h3v3h-1v7h1v3h-3v-1h-7v1H8v-3h1v-3H5v1H2v-3h1V6H2V3m16 9v-1h-3v3h1v3h-3v-1h-3v3h1v1h7v-1h1v-7h-1m-5-6V5H5v1H4v8h1v1h4v-3H8V9h3v1h3V6h-1m-2 6h-1v3h3v-1h1v-3h-3v1M3 5h1V4H3v1m11 0h1V4h-1v1m-5 6h1v-1H9v1m10 0h1v-1h-1v1M9 21h1v-1H9v1m10 0h1v-1h-1v1M3 16h1v-1H3v1m11 0h1v-1h-1v1Z"/>`),
		g.Group(children),
	)
}