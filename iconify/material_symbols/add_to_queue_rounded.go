package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToQueueRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12v2q0 .425.288.713T12 15q.425 0 .713-.288T13 14v-2h2q.425 0 .713-.288T16 11q0-.425-.288-.713T15 10h-2V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v2H9q-.425 0-.713.288T8 11q0 .425.288.713T9 12h2Zm-7 7q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4Z"/>`),
		g.Group(children),
	)
}