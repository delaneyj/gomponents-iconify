package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraphicEqRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17V7q0-.425.288-.713T8 6q.425 0 .713.288T9 7v10q0 .425-.288.713T8 18q-.425 0-.713-.288T7 17Zm4 4V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v18q0 .425-.288.713T12 22q-.425 0-.713-.288T11 21Zm-8-8v-2q0-.425.288-.713T4 10q.425 0 .713.288T5 11v2q0 .425-.288.713T4 14q-.425 0-.713-.288T3 13Zm12 4V7q0-.425.288-.713T16 6q.425 0 .713.288T17 7v10q0 .425-.288.713T16 18q-.425 0-.713-.288T15 17Zm4-4v-2q0-.425.288-.713T20 10q.425 0 .713.288T21 11v2q0 .425-.288.713T20 14q-.425 0-.713-.288T19 13Z"/>`),
		g.Group(children),
	)
}