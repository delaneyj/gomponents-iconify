package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeetingRoomRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1V4q0-.425.288-.713T6 3h7q.425 0 .713.288T14 4h4q.425 0 .713.288T19 5v14h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21h-2q-.425 0-.713-.288T17 20V6h-3v14q0 .425-.288.713T13 21H4Zm8-9q0-.425-.288-.713T11 11q-.425 0-.713.288T10 12q0 .425.288.713T11 13q.425 0 .713-.288T12 12Z"/>`),
		g.Group(children),
	)
}