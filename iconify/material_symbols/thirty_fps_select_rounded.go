package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirtyFpsSelectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 14q-.425 0-.713-.288T4 13q0-.425.288-.713T5 12h4v-2H5q-.425 0-.713-.288T4 9q0-.425.288-.713T5 8h4V6H5q-.425 0-.713-.288T4 5q0-.425.288-.713T5 4h4q.825 0 1.413.588T11 6v1.5q0 .65-.425 1.075T9.5 9q.65 0 1.075.425T11 10.5V12q0 .825-.588 1.413T9 14H5Zm13 0h-3q-.825 0-1.413-.588T13 12V6q0-.825.588-1.413T15 4h3q.825 0 1.413.588T20 6v6q0 .825-.588 1.413T18 14Zm0-2V6h-3v6h3ZM4 22q-.425 0-.713-.288T3 21v-3q0-.425.288-.713T4 17q.425 0 .713.288T5 18v3q0 .425-.288.713T4 22Zm4 0q-.425 0-.713-.288T7 21v-3q0-.425.288-.713T8 17q.425 0 .713.288T9 18v3q0 .425-.288.713T8 22Zm4 0q-.425 0-.713-.288T11 21v-3q0-.425.288-.713T12 17q.425 0 .713.288T13 18v3q0 .425-.288.713T12 22Zm4 0q-.425 0-.713-.288T15 21v-3q0-.425.288-.713T16 17h4q.425 0 .713.288T21 18v3q0 .425-.288.713T20 22h-4Z"/>`),
		g.Group(children),
	)
}