package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSquareSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3zm-3 15c0 .6-.4 1-1 1s-1-.4-1-1v-4h-4v4c0 .6-.4 1-1 1s-1-.4-1-1V7c0-.6.4-1 1-1s1 .4 1 1v4h4V7c0-.6.4-1 1-1s1 .4 1 1v10z"/>`),
		g.Group(children),
	)
}