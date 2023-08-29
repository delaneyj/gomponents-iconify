package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeilingLightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.76 11l2 4H7.24l2-4h5.52M13 4h-2v5H8l-4 8h16l-4-8h-3V4m1 14h-4a2 2 0 1 0 4 0Z"/>`),
		g.Group(children),
	)
}