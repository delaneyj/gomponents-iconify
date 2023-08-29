package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorSensor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-.825 0-1.413-.588T7 19v-2h5q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13H7V5q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5v14q0 .825-.588 1.413T15 21H9Zm-4-5q-.425 0-.713-.288T4 15q0-.425.288-.713T5 14h7q.425 0 .713.288T13 15q0 .425-.288.713T12 16H5Zm7-6q.425 0 .713-.288T13 9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9q0 .425.288.713T12 10Zm8-1q-.425 0-.713-.288T19 8V4q0-.425.288-.713T20 3q.425 0 .713.288T21 4v4q0 .425-.288.713T20 9Z"/>`),
		g.Group(children),
	)
}