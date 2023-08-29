package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorBackOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2V3h14v16h2v2H3Zm4-2h10V5H7v14Zm3-6q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11q-.425 0-.713.288T9 12q0 .425.288.713T10 13ZM7 5v14V5Z"/>`),
		g.Group(children),
	)
}