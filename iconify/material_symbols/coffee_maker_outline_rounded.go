package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMakerOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h13q.425 0 .713.288T20 3q0 .425-.288.713T19 4h-1v2q0 .425-.288.713T17 7H9q-.425 0-.713-.288T8 6V4H6v16h4.05q-.95-.675-1.5-1.713T8 16v-3q0-.825.588-1.413T10 11h6q.825 0 1.413.588T18 13v3q0 1.25-.55 2.288T15.95 20H19q.425 0 .713.288T20 21q0 .425-.288.713T19 22H6Zm7-3q1.25 0 2.125-.875T16 16v-3h-6v3q0 1.25.875 2.125T13 19Zm0-9q.425 0 .713-.288T14 9q0-.425-.288-.713T13 8q-.425 0-.713.288T12 9q0 .425.288.713T13 10Zm0 3Z"/>`),
		g.Group(children),
	)
}