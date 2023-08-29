package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 5.5a.5.5 0 0 0-.748-.434l-3.5 2a.5.5 0 0 0 0 .868l3.5 2A.5.5 0 0 0 7 9.5V8.076l3.252 1.858A.5.5 0 0 0 11 9.5v-4a.5.5 0 0 0-.748-.434L7 6.924V5.5Z"/>`),
		g.Group(children),
	)
}