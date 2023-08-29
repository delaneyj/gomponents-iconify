package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNotesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q.425 0 .713-.288T8 13q0-.425-.288-.713T7 12q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-3q.425 0 .713-.288T8 10q0-.425-.288-.713T7 9q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0-3q.425 0 .713-.288T8 7q0-.425-.288-.713T7 6q-.425 0-.713.288T6 7q0 .425.288.713T7 8Zm4 6h3q.425 0 .713-.288T15 13q0-.425-.288-.713T14 12h-3q-.425 0-.713.288T10 13q0 .425.288.713T11 14Zm0-3h6q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-6q-.425 0-.713.288T10 10q0 .425.288.713T11 11Zm0-3h6q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6h-6q-.425 0-.713.288T10 7q0 .425.288.713T11 8ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm-2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}