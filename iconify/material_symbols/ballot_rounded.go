package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 10q.425 0 .713-.288T17 9q0-.425-.288-.713T16 8h-3q-.425 0-.713.288T12 9q0 .425.288.713T13 10h3Zm0 6q.425 0 .713-.288T17 15q0-.425-.288-.713T16 14h-3q-.425 0-.713.288T12 15q0 .425.288.713T13 16h3Zm-7-5q.825 0 1.413-.588T11 9q0-.825-.588-1.413T9 7q-.825 0-1.413.588T7 9q0 .825.588 1.413T9 11Zm0 6q.825 0 1.413-.588T11 15q0-.825-.588-1.413T9 13q-.825 0-1.413.588T7 15q0 .825.588 1.413T9 17Zm-4 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}