package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM10 5h1v5h-1V5Zm-5.252.066A.5.5 0 0 0 4 5.5v4a.5.5 0 0 0 .748.434l3.5-2a.5.5 0 0 0 0-.868l-3.5-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}