package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h2q.425 0 .713.288T6 3v19q-.825 0-1.413-.588T4 20V4H3q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2Zm6 20q-.825 0-1.413-.588T7 20V4q0-.825.588-1.413T9 2h9q.825 0 1.413.588T20 4v16q0 .825-.588 1.413T18 22H9Zm0-11.475q.45-.275.95-.4T11 10h5q.55 0 1.05.125t.95.4V4H9v6.525ZM13.5 9q-.825 0-1.413-.588T11.5 7q0-.825.588-1.413T13.5 5q.825 0 1.413.588T15.5 7q0 .825-.588 1.413T13.5 9Zm-1 8v1q0 .425.288.713T13.5 19q.425 0 .713-.288T14.5 18v-1h1q.425 0 .713-.288T16.5 16q0-.425-.288-.713T15.5 15h-1v-1q0-.425-.288-.713T13.5 13q-.425 0-.713.288T12.5 14v1h-1q-.425 0-.713.288T10.5 16q0 .425.288.713T11.5 17h1Z"/>`),
		g.Group(children),
	)
}