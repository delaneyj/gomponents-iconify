package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNotesOffOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q-.425 0-.713-.288T6 13q0-.425.288-.713T7 12q.425 0 .713.288T8 13q0 .425-.288.713T7 14Zm13.7 3.85L18.85 16H20V4H6.85l-2-2H20q.825 0 1.413.588T22 4v12q0 .65-.363 1.137t-.937.713ZM13.85 11l-2-2H17q.425 0 .713.288T18 10q0 .425-.288.713T17 11h-3.15ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4.8l-.6-.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L15.15 18H6ZM4 6.8V16h9.2L4 6.8Zm8.85 3.2ZM8.6 11.4ZM7 11q-.425 0-.713-.288T6 10q0-.425.288-.713T7 9q.425 0 .713.288T8 10q0 .425-.288.713T7 11Zm3.85-3L10 7.15V7q0-.425.288-.713T11 6h6q.425 0 .713.288T18 7q0 .425-.288.713T17 8h-6.15Z"/>`),
		g.Group(children),
	)
}