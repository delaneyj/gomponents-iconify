package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperBoardOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h14q.825 0 1.413.588T20 5v2h1q.425 0 .713.288T22 8q0 .425-.288.713T21 9h-1v2h1q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-1v2h1q.425 0 .713.288T22 16q0 .425-.288.713T21 17h-1v2q0 .825-.588 1.413T18 21H4Zm0-2h14V5H4v14Zm3-2h3q.425 0 .713-.288T11 16v-2q0-.425-.288-.713T10 13H7q-.425 0-.713.288T6 14v2q0 .425.288.713T7 17Zm6-7h2q.425 0 .713-.288T16 9V8q0-.425-.288-.713T15 7h-2q-.425 0-.713.288T12 8v1q0 .425.288.713T13 10Zm-6 2h3q.425 0 .713-.288T11 11V8q0-.425-.288-.713T10 7H7q-.425 0-.713.288T6 8v3q0 .425.288.713T7 12Zm6 5h2q.425 0 .713-.288T16 16v-4q0-.425-.288-.713T15 11h-2q-.425 0-.713.288T12 12v4q0 .425.288.713T13 17ZM4 5v14V5Z"/>`),
		g.Group(children),
	)
}