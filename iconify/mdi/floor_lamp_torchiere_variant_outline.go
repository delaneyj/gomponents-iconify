package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampTorchiereVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.97 4l-1.29 3h-1.36l-1.29-3h3.94M17 2H7l3 7h4l3-7m-1.08 20L13 15v7h-2v-7l-2.92 7H5.92l5-12h2.16l5 12h-2.16Z"/>`),
		g.Group(children),
	)
}