package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h8q.425 0 .713.288T14 4q0 .425-.288.713T13 5H5v14h14v-8q0-.425.288-.713T20 10q.425 0 .713.288T21 11v8q0 .825-.588 1.413T19 21H5Zm4-4q-.425 0-.713-.288T8 16q0-.425.288-.713T9 15h6q.425 0 .713.288T16 16q0 .425-.288.713T15 17H9Zm0-3q-.425 0-.713-.288T8 13q0-.425.288-.713T9 12h6q.425 0 .713.288T16 13q0 .425-.288.713T15 14H9Zm0-3q-.425 0-.713-.288T8 10q0-.425.288-.713T9 9h6q.425 0 .713.288T16 10q0 .425-.288.713T15 11H9Zm9-2q-.425 0-.713-.288T17 8V7h-1q-.425 0-.713-.288T15 6q0-.425.288-.713T16 5h1V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v1h1q.425 0 .713.288T21 6q0 .425-.288.713T20 7h-1v1q0 .425-.288.713T18 9Z"/>`),
		g.Group(children),
	)
}