package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellphoneCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.54 23H7a2 2 0 0 1-2-2V3c0-1.11.89-2 2-2h10a2 2 0 0 1 2 2v10c-.7 0-1.37.13-2 .35V5H7v14h6c0 1.54.58 2.94 1.54 4m3.21-.84l-2.75-3L16.16 18l1.59 1.59L21.34 16l1.16 1.41l-4.75 4.75"/>`),
		g.Group(children),
	)
}