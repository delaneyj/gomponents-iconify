package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpenOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13q.425 0 .713-.288T12 12q0-.425-.288-.713T11 11q-.425 0-.713.288T10 12q0 .425.288.713T11 13Zm-4 8v-2l6-1V5.975L7 5V3l8 1.3v15.35L7 21Zm-4 0v-2h2V3h14v16h2v2H3Zm4-2h10V5H7v14Z"/>`),
		g.Group(children),
	)
}