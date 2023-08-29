package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7V4H9v3h2.868L9.012 17H5v3h10v-3h-2.868l2.856-10z"/>`),
		g.Group(children),
	)
}