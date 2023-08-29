package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllOutRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21H5q-.825 0-1.413-.588T3 19v-2q0-.425.288-.713T4 16q.425 0 .713.288T5 17v2h2q.425 0 .713.288T8 20q0 .425-.288.713T7 21Zm12 0h-2q-.425 0-.713-.288T16 20q0-.425.288-.713T17 19h2v-2q0-.425.288-.713T20 16q.425 0 .713.288T21 17v2q0 .825-.588 1.413T19 21Zm-7-2q-2.9 0-4.95-2.05T5 12q0-2.9 2.05-4.95T12 5q2.9 0 4.95 2.05T19 12q0 2.9-2.05 4.95T12 19ZM3 5q0-.825.588-1.413T5 3h2q.425 0 .713.288T8 4q0 .425-.288.713T7 5H5v2q0 .425-.288.713T4 8q-.425 0-.713-.288T3 7V5Zm17 3q-.425 0-.713-.288T19 7V5h-2q-.425 0-.713-.288T16 4q0-.425.288-.713T17 3h2q.825 0 1.413.588T21 5v2q0 .425-.288.713T20 8Z"/>`),
		g.Group(children),
	)
}