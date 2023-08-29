package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsGolfRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16q-2.925 0-4.963-2.038T5 9q0-2.925 2.038-4.963T12 2q2.925 0 4.963 2.038T19 9q0 2.925-2.038 4.963T12 16Zm-2-7q.425 0 .713-.288T11 8q0-.425-.288-.713T10 7q-.425 0-.713.288T9 8q0 .425.288.713T10 9Zm4 0q.425 0 .713-.288T15 8q0-.425-.288-.713T14 7q-.425 0-.713.288T13 8q0 .425.288.713T14 9Zm-2-2q.425 0 .713-.288T13 6q0-.425-.288-.713T12 5q-.425 0-.713.288T11 6q0 .425.288.713T12 7Zm-1 15v-1q0-.825-.588-1.413T9 19H8q-.425 0-.713-.288T7 18q0-.425.288-.713T8 17h8q.425 0 .713.288T17 18q0 .425-.288.713T16 19h-1q-.825 0-1.413.588T13 21v1h-2Z"/>`),
		g.Group(children),
	)
}