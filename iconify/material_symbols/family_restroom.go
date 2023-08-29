package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamilyRestroom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6q-.825 0-1.413-.588T16 4q0-.825.588-1.413T18 2q.825 0 1.413.588T20 4q0 .825-.588 1.413T18 6Zm-1 16v-8q0-1-.513-1.8t-1.312-1.25l.875-2.575q.2-.625.738-1T18 7q.675 0 1.213.375t.737 1L22.5 16H20v6h-3Zm-4.5-10.5q-.625 0-1.063-.438T11 10q0-.625.438-1.063T12.5 8.5q.625 0 1.063.438T14 10q0 .625-.438 1.063T12.5 11.5ZM5.5 6q-.825 0-1.413-.588T3.5 4q0-.825.588-1.413T5.5 2q.825 0 1.413.588T7.5 4q0 .825-.588 1.413T5.5 6Zm-2 16v-7H2V9q0-.825.588-1.413T4 7h3q.825 0 1.413.588T9 9v6H7.5v7h-4Zm7.5 0v-4h-1v-4q0-.625.438-1.063T11.5 12.5h2q.625 0 1.063.438T15 14v4h-1v4h-3Z"/>`),
		g.Group(children),
	)
}