package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4v10.273l-3 3.652l-3-3.652V4H7v10.273a2 2 0 0 0 .455 1.27l3.772 4.592h1.546l3.772-4.592a2 2 0 0 0 .455-1.27V4h-2Z"/>`),
		g.Group(children),
	)
}