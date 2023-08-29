package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNotesOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4.5h2.5L11 11H8.2L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L15.15 18H6Zm14.7-.15L13.85 11H17q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9h-5.15l-1-1H17q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6h-6q-.425 0-.713.288T10 7v.15L4.85 2H20q.825 0 1.413.588T22 4v12q0 .65-.363 1.137t-.937.713ZM7 14q.425 0 .713-.288T8 13q0-.425-.288-.713T7 12q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-3q.425 0 .713-.288T8 10q0-.425-.288-.713T7 9q-.425 0-.713.288T6 10q0 .425.288.713T7 11Z"/>`),
		g.Group(children),
	)
}