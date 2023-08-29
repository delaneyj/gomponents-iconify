package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermContactCalendarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 18.85q1.35-1.325 3.138-2.087T12 16q2.075 0 3.863.763T19 18.85V6H5v12.85ZM12 14q-1.45 0-2.475-1.025T8.5 10.5q0-1.45 1.025-2.475T12 7q1.45 0 2.475 1.025T15.5 10.5q0 1.45-1.025 2.475T12 14Zm-7 8q-.825 0-1.413-.588T3 20V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h8V3q0-.425.288-.713T17 2q.425 0 .713.288T18 3v1h1q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22H5Z"/>`),
		g.Group(children),
	)
}