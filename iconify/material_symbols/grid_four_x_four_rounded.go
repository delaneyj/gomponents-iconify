package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridFourXFourRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19H3q-.425 0-.713-.288T2 18q0-.425.288-.713T3 17h2v-4H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h2V7H3q-.425 0-.713-.288T2 6q0-.425.288-.713T3 5h2V3q0-.425.288-.713T6 2q.425 0 .713.288T7 3v2h4V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v2h4V3q0-.425.288-.713T18 2q.425 0 .713.288T19 3v2h2q.425 0 .713.288T22 6q0 .425-.288.713T21 7h-2v4h2q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-2v4h2q.425 0 .713.288T22 18q0 .425-.288.713T21 19h-2v2q0 .425-.288.713T18 22q-.425 0-.713-.288T17 21v-2h-4v2q0 .425-.288.713T12 22q-.425 0-.713-.288T11 21v-2H7v2q0 .425-.288.713T6 22q-.425 0-.713-.288T5 21v-2Zm2-2h4v-4H7v4Zm6 0h4v-4h-4v4Zm-6-6h4V7H7v4Zm6 0h4V7h-4v4Z"/>`),
		g.Group(children),
	)
}