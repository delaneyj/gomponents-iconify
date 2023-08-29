package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipExitOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18v-6q0-.425.288-.713T3 11q.425 0 .713.288T4 12v6h16V6h-8q-.425 0-.713-.288T11 5q0-.425.288-.713T12 4h8q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm10-6.575l2.375 2.375q.3.3.7.3t.7-.3q.3-.3.3-.713t-.3-.712L15.4 12H17q.425 0 .713-.288T18 11q0-.425-.288-.713T17 10h-4q-.425 0-.713.288T12 11v4q0 .425.288.713T13 16q.425 0 .713-.288T14 15v-1.575ZM3 9q-.425 0-.713-.288T2 8V5q0-.425.288-.713T3 4h5q.425 0 .713.288T9 5v3q0 .425-.288.713T8 9H3Zm9 3Z"/>`),
		g.Group(children),
	)
}