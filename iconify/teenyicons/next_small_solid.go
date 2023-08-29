package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.748 5.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2ZM10 10h1V5h-1v5Z"/>`),
		g.Group(children),
	)
}