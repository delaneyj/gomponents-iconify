package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptLong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-1.25 0-2.125-.875T3 19v-3h3V2l1.5 1.5L9 2l1.5 1.5L12 2l1.5 1.5L15 2l1.5 1.5L18 2l1.5 1.5L21 2v17q0 1.25-.875 2.125T18 22H6Zm12-2q.425 0 .713-.288T19 19V5H8v11h9v3q0 .425.288.713T18 20ZM9 9V7h6v2H9Zm0 3v-2h6v2H9Zm8-3q-.425 0-.713-.288T16 8q0-.425.288-.713T17 7q.425 0 .713.288T18 8q0 .425-.288.713T17 9Zm0 3q-.425 0-.713-.288T16 11q0-.425.288-.713T17 10q.425 0 .713.288T18 11q0 .425-.288.713T17 12Z"/>`),
		g.Group(children),
	)
}