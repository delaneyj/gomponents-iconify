package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterFocusWeakRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-3q0-.425.288-.713T4 15q.425 0 .713.288T5 16v3h3q.425 0 .713.288T9 20q0 .425-.288.713T8 21H5Zm14 0h-3q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19h3v-3q0-.425.288-.713T20 15q.425 0 .713.288T21 16v3q0 .825-.588 1.413T19 21ZM3 8V5q0-.825.588-1.413T5 3h3q.425 0 .713.288T9 4q0 .425-.288.713T8 5H5v3q0 .425-.288.713T4 9q-.425 0-.713-.288T3 8Zm16 0V5h-3q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h3q.825 0 1.413.588T21 5v3q0 .425-.288.713T20 9q-.425 0-.713-.288T19 8Zm-7 8q-1.65 0-2.825-1.175T8 12q0-1.65 1.175-2.825T12 8q1.65 0 2.825 1.175T16 12q0 1.65-1.175 2.825T12 16Z"/>`),
		g.Group(children),
	)
}