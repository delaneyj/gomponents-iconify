package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 14h1q.425 0 .713-.288T10.5 13v-1h-1v1h-1v-2h2q0-.425-.288-.713T9.5 10h-1q-.425 0-.713.288T7.5 11v2q0 .425.288.713T8.5 14Zm3 0h1v-4h-1v4Zm2 0h1v-1.5H16v-1h-1.5V11h2v-1h-3v4ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}