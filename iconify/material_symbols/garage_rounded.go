package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm5-8q-.425 0-.713-.288T8 13q0-.425.288-.713T9 12q.425 0 .713.288T10 13q0 .425-.288.713T9 14Zm6 0q-.425 0-.713-.288T14 13q0-.425.288-.713T15 12q.425 0 .713.288T16 13q0 .425-.288.713T15 14Zm-9 4.5q.425 0 .713-.288T7 17.5v-1h10v1q0 .425.288.713T18 18.5q.425 0 .713-.288T19 17.5v-6.075q0-.175-.025-.338t-.075-.312L17.55 6.85q-.2-.6-.725-.975T15.65 5.5h-7.3q-.65 0-1.175.375t-.725.975L5.1 10.775q-.05.15-.075.313T5 11.425V17.5q0 .425.287.713T6 18.5Zm1.65-9l.7-2h7.3l.7 2h-8.7Z"/>`),
		g.Group(children),
	)
}