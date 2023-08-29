package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RavenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.35 22l-1.85-.75l1.45-3.525q-2.65-.7-4.3-2.85T2 10V6q0-1.65 1.175-2.825T6 2q.55 0 1.05.187t1 .388L14 5l-4 1.5V8l11 7l1 5h-2l-1-2h-5v4h-2v-4h-2l-1.65 4ZM10 16h8.825l-1.575-1H10q-1.65 0-2.825-1.175T6 11h2q0 .825.588 1.413T10 13h4.125L8 9.1V6q0-.825-.588-1.413T6 4q-.825 0-1.413.588T4 6v4q0 2.5 1.75 4.25T10 16ZM6 7q-.425 0-.713-.288T5 6q0-.425.288-.713T6 5q.425 0 .713.288T7 6q0 .425-.288.713T6 7Zm4 8Z"/>`),
		g.Group(children),
	)
}