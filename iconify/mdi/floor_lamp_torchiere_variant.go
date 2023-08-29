package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorLampTorchiereVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 2l3 7h4l3-7H7m8.92 20L13 15v7h-2v-7l-2.92 7H5.92l5-12h2.16l5 12h-2.16Z"/>`),
		g.Group(children),
	)
}