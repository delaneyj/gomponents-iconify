package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToVoteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.175 16h13.65l-1.95-2.2l1.425-1.425L21 15.45V22H3v-6.55l2.75-3.125l1.425 1.425l-2 2.25Zm6.85-.175L5.7 9.45l7.75-7.75l6.375 6.325l-7.8 7.8Zm.025-2.875L17 8l-3.55-3.5L8.5 9.45l3.55 3.5Z"/>`),
		g.Group(children),
	)
}