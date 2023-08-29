package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToVoteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.175 16h13.65l-1.95-2.2l1.425-1.425L21 15.45V22H3v-6.55l2.75-3.125l1.425 1.425l-2 2.25Zm6.875-.2L5.7 9.45l7.775-7.775l6.35 6.35L12.05 15.8Zm0-2.825L17 8.025L13.475 4.5l-4.95 4.95l3.525 3.525ZM5 20h14v-2H5v2Zm0 0v-2v2Z"/>`),
		g.Group(children),
	)
}