package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatasetLinkedOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v7.1q-.25-.05-.488-.075T20 11h-1V4H5v14h3.1q.1.55.263 1.05T8.8 20H5Zm0-3v1V4v13Zm2-1h1.1q.2-1.225.875-2.263T10.7 12H7v4Zm0-6h4V6H7v4Zm7 11q-1.65 0-2.825-1.175T10 17q0-1.65 1.175-2.825T14 13h1q.425 0 .713.288T16 14q0 .425-.288.713T15 15h-1q-.825 0-1.413.588T12 17q0 .825.588 1.413T14 19h1q.425 0 .713.288T16 20q0 .425-.288.713T15 21h-1Zm-1-11h4V6h-4v4Zm2 8q-.425 0-.713-.288T14 17q0-.425.288-.713T15 16h4q.425 0 .713.288T20 17q0 .425-.288.713T19 18h-4Zm5 3h-1q-.425 0-.713-.288T18 20q0-.425.288-.713T19 19h1q.825 0 1.413-.588T22 17q0-.825-.588-1.413T20 15h-1q-.425 0-.713-.288T18 14q0-.425.288-.713T19 13h1q1.65 0 2.825 1.163T24 17q0 1.65-1.175 2.825T20 21Z"/>`),
		g.Group(children),
	)
}