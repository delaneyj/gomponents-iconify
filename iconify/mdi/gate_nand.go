package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateNand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4v16h8c3.43 0 6.5-2.16 7.6-5.4c.4.2.9.4 1.4.4a3 3 0 0 0 3-3a3 3 0 0 0-3-3c-.5 0-.97.15-1.4.4C16.5 6.16 13.43 4 10 4H2m2 2h6a6 6 0 0 1 6 6a6 6 0 0 1-6 6H4V6m15 5c.5 0 1 .5 1 1s-.5 1-1 1a1 1 0 0 1-1-1c0-.5.5-1 1-1Z"/>`),
		g.Group(children),
	)
}