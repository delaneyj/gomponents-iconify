package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowShutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h18v4h-2v12h-2V8H7v12H5V8H3V4m5 5h8v2H8V9m0 3h8v2H8v-2m0 3h8v2H8v-2m0 3h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}