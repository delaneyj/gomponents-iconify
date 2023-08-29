package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorFrontOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2V3h14v16h2v2H3Zm4-2h10V5H7v14Zm7-6q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11q-.425 0-.713.288T13 12q0 .425.288.713T14 13ZM7 5v14V5Z"/>`),
		g.Group(children),
	)
}