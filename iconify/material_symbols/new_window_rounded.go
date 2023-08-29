package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewWindowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h5q.425 0 .713.288T11 4q0 .425-.288.713T10 5H5v14h14v-5q0-.425.288-.713T20 13q.425 0 .713.288T21 14v5q0 .825-.588 1.413T19 21H5ZM16 8h-2q-.425 0-.713-.288T13 7q0-.425.288-.713T14 6h2V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v2h2q.425 0 .713.288T21 7q0 .425-.288.713T20 8h-2v2q0 .425-.288.713T17 11q-.425 0-.713-.288T16 10V8Z"/>`),
		g.Group(children),
	)
}