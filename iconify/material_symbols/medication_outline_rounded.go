package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.5q.625 0 1.063-.438T13.5 16v-1h1q.625 0 1.063-.438T16 13.5q0-.625-.438-1.063T14.5 12h-1v-1q0-.625-.438-1.063T12 9.5q-.625 0-1.063.438T10.5 11v1h-1q-.625 0-1.063.438T8 13.5q0 .625.438 1.063T9.5 15h1v1q0 .625.438 1.063T12 17.5ZM7 21q-.825 0-1.413-.588T5 19V8q0-.825.588-1.413T7 6h10q.825 0 1.413.588T19 8v11q0 .825-.588 1.413T17 21H7Zm0-2h10V8H7v11ZM7 5q-.425 0-.713-.288T6 4q0-.425.288-.713T7 3h10q.425 0 .713.288T18 4q0 .425-.288.713T17 5H7Zm0 3v11V8Z"/>`),
		g.Group(children),
	)
}