package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Print(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 0v2h4V0H2zM.09 3C.03 3 0 3.04 0 3.09V5.9c0 .05.04.09.09.09H1v-2h6v2h.91c.05 0 .09-.04.09-.09V3.09C8 3.03 7.96 3 7.91 3H.1zM2 5v3h4V5H2z"/>`),
		g.Group(children),
	)
}