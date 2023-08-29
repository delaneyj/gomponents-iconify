package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightTrack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2L5 5v11l1 4c.27 1.07.9 2 2 2h8a2 2 0 0 0 2-2l1-4V5l-2-3h-2v1h-2V2H7m0 4h10v10H7V6Z"/>`),
		g.Group(children),
	)
}