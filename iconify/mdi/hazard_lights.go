package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HazardLights(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 12l2.33 4H9.68L12 12m0-4L6.21 18H17.8L12 8m0-6L1 21h22L12 2m0 4l7.53 13H4.47L12 6Z"/>`),
		g.Group(children),
	)
}