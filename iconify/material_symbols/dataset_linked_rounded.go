package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatasetLinkedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v7.075q-.25-.05-.488-.063T20 11h-6q-.925 0-1.763.263T10.676 12H7v4h1.075q-.05.25-.063.488T8 17q0 .8.2 1.563T8.8 20H5Zm2-10h4V6H7v4Zm6 0h4V6h-4v4Zm1 11q-1.65 0-2.825-1.175T10 17q0-1.65 1.175-2.825T14 13h1q.425 0 .713.288T16 14q0 .425-.288.713T15 15h-1q-.825 0-1.413.588T12 17q0 .825.588 1.413T14 19h1q.425 0 .713.288T16 20q0 .425-.288.713T15 21h-1Zm1-3q-.425 0-.713-.288T14 17q0-.425.288-.713T15 16h4q.425 0 .713.288T20 17q0 .425-.288.713T19 18h-4Zm5 3h-1q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19h1q.825 0 1.413-.588T22 17q0-.825-.588-1.413T20 15h-1q-.425 0-.713-.288T18 14q0-.425.288-.713T19 13h1q1.65 0 2.825 1.163T24 17q0 1.65-1.175 2.825T20 21Z"/>`),
		g.Group(children),
	)
}