package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabProfileRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12h6q.425 0 .713-.288T16 11q0-.425-.288-.713T15 10H9q-.425 0-.713.288T8 11q0 .425.288.713T9 12Zm0-4h6q.425 0 .713-.288T16 7q0-.425-.288-.713T15 6H9q-.425 0-.713.288T8 7q0 .425.288.713T9 8Zm10.95 12.475L15.9 15.2q-.425-.575-1.05-.887T13.5 14H4V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v16q0 .125-.013.238t-.037.237ZM6 22q-.825 0-1.412-.588T4 20v-4h9.5q.25 0 .463.113t.362.312l4.2 5.5q-.125.05-.262.063T18 22H6Z"/>`),
		g.Group(children),
	)
}