package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRepairServiceOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14ZM2 19v-9q0-.825.588-1.413T4 8h3V6q0-.825.588-1.413T9 4h6q.825 0 1.413.588T17 6v2h3q.825 0 1.413.588T22 10v9q0 .425-.288.713T21 20H3q-.425 0-.713-.288T2 19Zm6-4q0 .425-.288.713T7 16q-.425 0-.713-.288T6 15H4v3h16v-3h-2q0 .425-.288.713T17 16q-.425 0-.713-.288T16 15H8Zm-4-5v3h2q0-.425.288-.713T7 12q.425 0 .713.288T8 13h8q0-.425.288-.713T17 12q.425 0 .713.288T18 13h2v-3H4Zm5-2h6V6H9v2Z"/>`),
		g.Group(children),
	)
}