package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm6.249-2.432a.5.5 0 0 1 .5-.002l3.5 2a.5.5 0 0 1 0 .868l-3.5 2A.5.5 0 0 1 6 9.5v-4a.5.5 0 0 1 .249-.432Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}