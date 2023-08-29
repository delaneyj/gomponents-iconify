package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QueuePlayNextRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12v2q0 .425.288.713T12 15q.425 0 .713-.288T13 14v-2h2q.425 0 .713-.288T16 11q0-.425-.288-.713T15 10h-2V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v2H9q-.425 0-.713.288T8 11q0 .425.288.713T9 12h2Zm10 6l-2.25-2.25q-.325-.325-.325-.75t.325-.75q.325-.325.75-.325t.75.325l3.05 3.05q.3.3.3.7t-.3.7l-3.05 3.05q-.325.325-.75.325t-.75-.325q-.325-.325-.325-.75t.325-.75L21 18ZM9 21q-.425 0-.713-.288T8 20v-1H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v6q0 .425-.288.713T21 12h-2q-.825 0-1.413.588T17 14v4q0 .425-.288.713T16 19h-1v1q0 .425-.288.713T14 21H9Z"/>`),
		g.Group(children),
	)
}