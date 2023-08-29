package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Powerpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3H6zm2 2h16v22H8V5zm5 6v2h4c1.19 0 2 .81 2 2s-.81 2-2 2s-2-.81-2-2h-2v7h2v-3.594c.594.35 1.26.594 2 .594c2.21 0 4-1.79 4-4s-1.79-4-4-4h-4z"/>`),
		g.Group(children),
	)
}