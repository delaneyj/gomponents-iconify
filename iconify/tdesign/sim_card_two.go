package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586L7.586 1Zm.828 2L5 6.414V21h14V3H8.414ZM12 10a1 1 0 0 0-1 1v1H9v-1a3 3 0 1 1 5.149 2.094l-.028.028l-1.841 1.611H15v2H9v-1.787l3.739-3.272A1 1 0 0 0 12 10Z"/>`),
		g.Group(children),
	)
}