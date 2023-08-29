package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvenGenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-9v7h14v-7h-2v5H7v-5H5Zm4 3h6v-3H9v3Zm-4-5h14V5H5v5Zm3-2q-.425 0-.713-.288T7 7q0-.425.288-.713T8 6q.425 0 .713.288T9 7q0 .425-.288.713T8 8Zm4 0q-.425 0-.713-.288T11 7q0-.425.288-.713T12 6q.425 0 .713.288T13 7q0 .425-.288.713T12 8Zm4 0q-.425 0-.713-.288T15 7q0-.425.288-.713T16 6q.425 0 .713.288T17 7q0 .425-.288.713T16 8Z"/>`),
		g.Group(children),
	)
}