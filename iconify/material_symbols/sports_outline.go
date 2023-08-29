package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19q-2.5 0-4.25-1.75T5 13q0-.275.025-.55t.075-.55q-.125.05-.3.075T4.5 12q-1.05 0-1.775-.725T2 9.5q0-1.05.688-1.775T4.425 7q.825 0 1.488.462T6.85 8.65q.825-.75 1.888-1.2T11 7h11v4h-5v2q0 2.5-1.75 4.25T11 19Zm-6.5-8.5q.425 0 .713-.288T5.5 9.5q0-.425-.288-.713T4.5 8.5q-.425 0-.713.288T3.5 9.5q0 .425.288.713t.712.287Zm6.5 6q1.45 0 2.475-1.025T14.5 13q0-1.45-1.025-2.475T11 9.5q-1.45 0-2.475 1.025T7.5 13q0 1.45 1.025 2.475T11 16.5Zm0-1.5q.825 0 1.413-.588T13 13q0-.825-.588-1.413T11 11q-.825 0-1.413.588T9 13q0 .825.588 1.413T11 15Zm0-2Z"/>`),
		g.Group(children),
	)
}