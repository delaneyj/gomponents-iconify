package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M.09 0C.03 0 0 .04 0 .09V7.9c0 .05.04.09.09.09H7.9c.05 0 .09-.04.09-.09V.09C7.99.03 7.95 0 7.9 0H.09zM1.5.78L3.22 2.5L1.5 4.22L.78 3.5l1-1l-1-1L1.5.78zM4 3h3v1H4V3z"/>`),
		g.Group(children),
	)
}