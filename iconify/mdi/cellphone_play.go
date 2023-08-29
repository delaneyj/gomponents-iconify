package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphonePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19H7V5h10m0-4H7c-1.11 0-2 .89-2 2v18c0 1.1.9 2 2 2h10c1.1 0 2-.9 2-2V3a2 2 0 0 0-2-2m-7 8v6l4-3l-4-3Z"/>`),
		g.Group(children),
	)
}