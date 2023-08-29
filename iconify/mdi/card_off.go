package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.88 18.68L7.2 4H20c1.11 0 2 .89 2 2v12c0 .24-.04.47-.12.68m-1.32 1.23h.01L2.39 1.73L1.11 3l1.54 1.54C2.25 4.9 2 5.42 2 6v12a2 2 0 0 0 2 2h14.11l2.73 2.73l1.27-1.27l-1.55-1.55Z"/>`),
		g.Group(children),
	)
}