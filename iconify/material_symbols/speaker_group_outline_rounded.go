package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerGroupOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 17h9V3h-9v14Zm0 2q-.825 0-1.413-.588T8 17V3q0-.825.588-1.413T10 1h9q.825 0 1.413.588T21 3v14q0 .825-.588 1.413T19 19h-9Zm4.5-11.5q.625 0 1.063-.438T16 6q0-.625-.438-1.063T14.5 4.5q-.625 0-1.063.438T13 6q0 .625.438 1.063T14.5 7.5Zm0 8.5q1.45 0 2.475-1.025T18 12.5q0-1.45-1.025-2.475T14.5 9q-1.45 0-2.475 1.025T11 12.5q0 1.45 1.025 2.475T14.5 16Zm0-2q-.625 0-1.063-.438T13 12.5q0-.625.438-1.063T14.5 11q.625 0 1.063.438T16 12.5q0 .625-.438 1.063T14.5 14ZM6 23q-.825 0-1.413-.588T4 21V6q0-.425.288-.713T5 5q.425 0 .713.288T6 6v15h9q.425 0 .713.288T16 22q0 .425-.288.713T15 23H6Zm4-20v14V3Z"/>`),
		g.Group(children),
	)
}