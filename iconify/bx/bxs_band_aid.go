package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBandAid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M11.646 3.868l-7.778 7.778a6.007 6.007 0 0 0 0 8.485a5.984 5.984 0 0 0 4.242 1.754a5.984 5.984 0 0 0 4.243-1.754l7.778-7.778a6.007 6.007 0 0 0 0-8.485a6.008 6.008 0 0 0-8.485 0zM9 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-6a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}