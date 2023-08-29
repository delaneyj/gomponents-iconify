package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamilyRestroomSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6q-.825 0-1.413-.588T16 4q0-.825.588-1.413T18 2q.825 0 1.413.588T20 4q0 .825-.588 1.413T18 6Zm-1 16V10.95h-1.825L16.525 7h2.95l3.025 9H20v6h-3Zm-4.5-10.5q-.625 0-1.063-.438T11 10q0-.625.438-1.063T12.5 8.5q.625 0 1.063.438T14 10q0 .625-.438 1.063T12.5 11.5ZM5.5 6q-.825 0-1.413-.588T3.5 4q0-.825.588-1.413T5.5 2q.825 0 1.413.588T7.5 4q0 .825-.588 1.413T5.5 6Zm-2 16v-7H2V7h7v8H7.5v7h-4Zm7.5 0v-4h-1v-5.5h5V18h-1v4h-3Z"/>`),
		g.Group(children),
	)
}