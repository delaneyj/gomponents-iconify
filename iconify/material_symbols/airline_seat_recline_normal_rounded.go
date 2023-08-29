package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatReclineNormalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20H8q-.825 0-1.413-.588T6 18V8q0-.425.288-.713T7 7q.425 0 .713.288T8 8v10h6q.425 0 .713.288T15 19q0 .425-.288.713T14 20ZM11.5 6q-.825 0-1.413-.588T9.5 4q0-.825.588-1.413T11.5 2q.825 0 1.413.588T13.5 4q0 .825-.588 1.413T11.5 6ZM16 21v-4h-5q-.825 0-1.413-.588T9 15V9.5q0-1.05.725-1.775T11.5 7q1.05 0 1.775.725T14 9.5V14h2q.825 0 1.413.588T18 16v5q0 .425-.288.713T17 22q-.425 0-.713-.288T16 21Z"/>`),
		g.Group(children),
	)
}