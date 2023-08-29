package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventNoteRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h8V3q0-.425.288-.713T17 2q.425 0 .713.288T18 3v1h1q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22H5Zm0-2h14V10H5v10Zm3-6q-.425 0-.713-.288T7 13q0-.425.288-.713T8 12h8q.425 0 .713.288T17 13q0 .425-.288.713T16 14H8Zm0 4q-.425 0-.713-.288T7 17q0-.425.288-.713T8 16h5q.425 0 .713.288T14 17q0 .425-.288.713T13 18H8Z"/>`),
		g.Group(children),
	)
}