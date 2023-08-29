package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5q0 .425-.288.713T21 12q-.425 0-.713-.288T20 11V6H4v12h9q.425 0 .713.288T14 19q0 .425-.288.713T13 20H4Zm8.4-10l2.25 2.25q.275.275.275.688t-.275.712q-.3.3-.713.3t-.712-.3L11 11.425v1.225q0 .425-.287.713T10 13.65q-.425 0-.713-.288T9 12.65V9q0-.425.288-.713T10 8h3.65q.425 0 .713.288T14.65 9q0 .425-.288.713T13.65 10H12.4ZM19 20q-1.25 0-2.125-.875T16 17q0-1.25.875-2.125T19 14q1.25 0 2.125.875T22 17q0 1.25-.875 2.125T19 20Z"/>`),
		g.Group(children),
	)
}