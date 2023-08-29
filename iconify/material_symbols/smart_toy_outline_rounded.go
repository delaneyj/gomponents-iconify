package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartToyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15q-1.25 0-2.125-.875T1 12q0-1.25.875-2.125T4 9V7q0-.825.588-1.413T6 5h3q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5h3q.825 0 1.413.588T20 7v2q1.25 0 2.125.875T23 12q0 1.25-.875 2.125T20 15v4q0 .825-.588 1.413T18 21H6q-.825 0-1.413-.588T4 19v-4Zm5-2q.625 0 1.063-.438T10.5 11.5q0-.625-.438-1.063T9 10q-.625 0-1.063.438T7.5 11.5q0 .625.438 1.063T9 13Zm6 0q.625 0 1.063-.438T16.5 11.5q0-.625-.438-1.063T15 10q-.625 0-1.063.438T13.5 11.5q0 .625.438 1.063T15 13Zm-6 4h6q.425 0 .713-.288T16 16q0-.425-.288-.713T15 15H9q-.425 0-.713.288T8 16q0 .425.288.713T9 17Zm-3 2h12V7H6v12Zm6-6Z"/>`),
		g.Group(children),
	)
}