package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewWeek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.75 19q-.425 0-.713-.288T12.75 18V6q0-.425.288-.713T13.75 5h1.375q.425 0 .713.288t.287.712v12q0 .425-.288.713t-.712.287H13.75Zm-4.875 0q-.425 0-.713-.288T7.875 18V6q0-.425.288-.713T8.874 5h1.375q.425 0 .713.288T11.25 6v12q0 .425-.288.713T10.25 19H8.875ZM4 19q-.425 0-.713-.288T3 18V6q0-.425.288-.713T4 5h1.375q.425 0 .713.288T6.375 6v12q0 .425-.288.713T5.376 19H4Zm14.625 0q-.425 0-.713-.288T17.626 18V6q0-.425.288-.713T18.625 5H20q.425 0 .713.288T21 6v12q0 .425-.288.713T20 19h-1.375Z"/>`),
		g.Group(children),
	)
}