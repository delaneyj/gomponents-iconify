package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 0-15 0a7.5 7.5 0 0 0 15 0Zm-4.249-2.432a.5.5 0 0 0-.5-.002L7 6.924V5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.249-.432Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}