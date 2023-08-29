package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCardOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6H4v6h9q.425 0 .713.288T14 19q0 .425-.288.713T13 20H4ZM4 8h16V6H4v2Zm15 11h-2q-.425 0-.713-.288T16 18q0-.425.288-.713T17 17h2v-2q0-.425.288-.713T20 14q.425 0 .713.288T21 15v2h2q.425 0 .713.288T24 18q0 .425-.288.713T23 19h-2v2q0 .425-.288.713T20 22q-.425 0-.713-.288T19 21v-2ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}