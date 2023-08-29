package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3.503a1.5 1.5 0 0 0-1.98-1.421l-11 3.715A1.5 1.5 0 0 0 1 7.217v1.566a1.5 1.5 0 0 0 1.018 1.42L4 10.877V12a3 3 0 0 0 5.872.868l3.146 1.067A1.5 1.5 0 0 0 15 12.515V3.503ZM5 11.215l3.924 1.332A2 2 0 0 1 5 12v-.785Z"/>`),
		g.Group(children),
	)
}