package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRepairServiceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8h6V6H9v2ZM3 20q-.425 0-.713-.288T2 19v-4h4q0 .425.288.713T7 16q.425 0 .713-.288T8 15h8q0 .425.288.713T17 16q.425 0 .713-.288T18 15h4v4q0 .425-.288.713T21 20H3Zm-1-6v-4q0-.825.588-1.413T4 8h3V6q0-.825.588-1.413T9 4h6q.825 0 1.413.588T17 6v2h3q.825 0 1.413.588T22 10v4h-4v-1q0-.425-.288-.713T17 12q-.425 0-.713.288T16 13v1H8v-1q0-.425-.288-.713T7 12q-.425 0-.713.288T6 13v1H2Z"/>`),
		g.Group(children),
	)
}