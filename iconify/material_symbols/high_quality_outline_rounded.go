package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighQualityOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.75 15v.75q0 .325.213.537t.537.213q.325 0 .537-.213t.213-.537V15H17q.425 0 .713-.288T18 14v-4q0-.425-.288-.713T17 9h-3q-.425 0-.713.288T13 10v4q0 .425.288.713T14 15h.75ZM7.5 13h2v1.25q0 .325.213.537t.537.213q.325 0 .537-.213T11 14.25v-4.5q0-.325-.213-.537T10.25 9q-.325 0-.537.213T9.5 9.75v1.75h-2V9.75q0-.325-.213-.537T6.75 9q-.325 0-.537.213T6 9.75v4.5q0 .325.213.537T6.75 15q.325 0 .537-.213t.213-.537V13Zm7 .5v-3h2v3h-2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}