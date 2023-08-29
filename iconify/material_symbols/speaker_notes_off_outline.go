package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNotesOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q-.425 0-.713-.288T6 13q0-.425.288-.713T7 12q.425 0 .713.288T8 13q0 .425-.288.713T7 14Zm13.7 3.85L18.85 16H20V4H6.85l-2-2H20q.825 0 1.413.588T22 4v12q0 .65-.363 1.137t-.937.713ZM13.85 11l-2-2H18v2h-4.15Zm6.65 12.3L15.15 18H6l-4 4V4.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM8.6 11.4Zm4.25-1.4ZM7 11q-.425 0-.713-.288T6 10q0-.425.288-.713T7 9q.425 0 .713.288T8 10q0 .425-.288.713T7 11Zm3.85-3L10 7.15V6h8v2h-7.15ZM4 6.8v10.325L5.15 16h8.05L4 6.8Z"/>`),
		g.Group(children),
	)
}