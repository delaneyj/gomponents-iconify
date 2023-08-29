package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarClutch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 18.84l4 1.86V23l-6-3v-6H5v-4h3V4l6-3v2.3l-4 1.86v13.68M19 10h-4V5.41L12 6.8v10.4l3 1.4V14h4v-4Z"/>`),
		g.Group(children),
	)
}