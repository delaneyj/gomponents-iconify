package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 8l3 5.2l3-2.7l-.7 3.5H6.7L6 10.5l3 2.7L12 8m0-4l-3.5 6L3 5l2 11h14l2-11l-5.5 5L12 4m7 14H5v1c0 .6.4 1 1 1h12c.6 0 1-.4 1-1v-1Z"/>`),
		g.Group(children),
	)
}