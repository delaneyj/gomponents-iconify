package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DvrRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14q.425 0 .713-.288T8 13q0-.425-.288-.713T7 12q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-4q.425 0 .713-.288T8 9q0-.425-.288-.713T7 8q-.425 0-.713.288T6 9q0 .425.288.713T7 10Zm3 4h7q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12h-7q-.425 0-.713.288T9 13q0 .425.288.713T10 14Zm0-4h7q.425 0 .713-.288T18 9q0-.425-.288-.713T17 8h-7q-.425 0-.713.288T9 9q0 .425.288.713T10 10Zm-6 9q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4Z"/>`),
		g.Group(children),
	)
}