package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 8l7 4.04V8l7 4V3h5v19H2V8m15 6l-7-4.27v4.04L3 9.73V21h17V4h-3v10M5 15h3v1H5v-1m0 3h5v1H5v-1m7-3h3v1h-3v-1m0 3h6v1h-6v-1Z"/>`),
		g.Group(children),
	)
}