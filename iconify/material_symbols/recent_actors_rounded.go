package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentActorsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h10q.825 0 1.413.588T15 7v10q0 .825-.588 1.413T13 19H3Zm0-3.65q1.1-.65 2.35-1T8 14q1.4 0 2.65.35t2.35 1V7H3v8.35Zm5-2.1q-1.15 0-1.95-.8t-.8-1.95q0-1.15.8-1.95T8 7.75q1.15 0 1.95.8t.8 1.95q0 1.15-.8 1.95t-1.95.8ZM18 19q-.425 0-.713-.288T17 18V6q0-.425.288-.713T18 5q.425 0 .713.288T19 6v12q0 .425-.288.713T18 19Zm4 0q-.425 0-.713-.288T21 18V6q0-.425.288-.713T22 5q.425 0 .713.288T23 6v12q0 .425-.288.713T22 19Z"/>`),
		g.Group(children),
	)
}