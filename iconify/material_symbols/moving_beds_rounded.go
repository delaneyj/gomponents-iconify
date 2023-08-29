package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovingBedsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h9q.825 0 1.413.588T15 4v16q0 .825-.588 1.413T13 22H4Zm0-11.475q.45-.275.95-.4T6 10h5q.55 0 1.05.125t.95.4V4H4v6.525ZM8.5 9q-.825 0-1.413-.588T6.5 7q0-.825.588-1.413T8.5 5q.825 0 1.413.588T10.5 7q0 .825-.588 1.413T8.5 9Zm-1 8v1q0 .425.288.713T8.5 19q.425 0 .713-.288T9.5 18v-1h1q.425 0 .713-.288T11.5 16q0-.425-.288-.713T10.5 15h-1v-1q0-.425-.288-.713T8.5 13q-.425 0-.713.288T7.5 14v1h-1q-.425 0-.713.288T5.5 16q0 .425.288.713T6.5 17h1Zm11.3-2.2q-.3-.3-.3-.713t.3-.712l.375-.375H17q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11h2.175l-.375-.375q-.3-.3-.3-.713t.3-.712q.3-.3.7-.3t.7.3l2.1 2.1q.3.3.3.7t-.3.7l-2.1 2.1q-.3.3-.7.3t-.7-.3Z"/>`),
		g.Group(children),
	)
}