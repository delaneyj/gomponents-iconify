package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorSlidingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13q-.425 0-.713-.288T8 12q0-.425.288-.713T9 11q.425 0 .713.288T10 12q0 .425-.288.713T9 13Zm6 0q-.425 0-.713-.288T14 12q0-.425.288-.713T15 11q.425 0 .713.288T16 12q0 .425-.288.713T15 13ZM4 21q-.425 0-.713-.288T3 20q0-.4.363-.563T4 19V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm2-2h5V5H6v14Zm7 0h5V5h-5v14Zm-1-8Z"/>`),
		g.Group(children),
	)
}