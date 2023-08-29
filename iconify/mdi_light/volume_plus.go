package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9h4l4-4h2v15h-2l-4-4H2V9m1 6h3.41l4 4H11V6h-.59l-4 4H3v5m14 1v-3h-3v-1h3V9h1v3h3v1h-3v3h-1Z"/>`),
		g.Group(children),
	)
}