package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KettleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17V6L4.2 3.6q-.375-.5-.1-1.05T5 2h10.775q.925 0 1.575.65T18 4.225V5h2q.825 0 1.413.588T22 7v5q0 .825-.588 1.413T20 14h-2v3q0 .825-.588 1.413T16 19H8q-.825 0-1.413-.588T6 17Zm12-5h2V7h-2v5Zm-4.5-7q-.625 0-1.063.438T12 6.5v8q0 .625.438 1.063T13.5 16q.625 0 1.063-.438T15 14.5v-8q0-.625-.438-1.063T13.5 5ZM4 22q-.425 0-.713-.288T3 21q0-.425.288-.713T4 20h16q.425 0 .713.288T21 21q0 .425-.288.713T20 22H4Z"/>`),
		g.Group(children),
	)
}