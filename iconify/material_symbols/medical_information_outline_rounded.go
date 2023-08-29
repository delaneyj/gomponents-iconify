package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalInformationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16v1q0 .425.288.713T8 18q.425 0 .713-.288T9 17v-1h1q.425 0 .713-.288T11 15q0-.425-.288-.713T10 14H9v-1q0-.425-.288-.713T8 12q-.425 0-.713.288T7 13v1H6q-.425 0-.713.288T5 15q0 .425.288.713T6 16h1Zm6.75-1.5h4.5q.325 0 .537-.213T19 13.75q0-.325-.213-.537T18.25 13h-4.5q-.325 0-.537.213T13 13.75q0 .325.213.537t.537.213Zm0 3h2.5q.325 0 .537-.213T17 16.75q0-.325-.213-.537T16.25 16h-2.5q-.325 0-.537.213T13 16.75q0 .325.213.537t.537.213ZM4 22q-.825 0-1.413-.588T2 20V9q0-.825.588-1.413T4 7h5V4q0-.825.588-1.413T11 2h2q.825 0 1.413.588T15 4v3h5q.825 0 1.413.588T22 9v11q0 .825-.588 1.413T20 22H4Zm0-2h16V9h-5q0 .825-.588 1.413T13 11h-2q-.825 0-1.413-.588T9 9H4v11Zm7-11h2V4h-2v5Zm1 5.5Z"/>`),
		g.Group(children),
	)
}