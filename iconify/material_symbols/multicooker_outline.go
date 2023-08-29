package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MulticookerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v1h3q.825 0 1.413.588T21 7v12q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19V7q0-.825.588-1.413T5 5h3ZM5 19h14v-9h-2v2q0 .825-.588 1.413T15 14H9q-.825 0-1.413-.588T7 12v-2H5v9Zm3-1q-.425 0-.713-.288T7 17q0-.425.288-.713T8 16q.425 0 .713.288T9 17q0 .425-.288.713T8 18Zm4 0q-.425 0-.713-.288T11 17q0-.425.288-.713T12 16q.425 0 .713.288T13 17q0 .425-.288.713T12 18Zm4 0q-.425 0-.713-.288T15 17q0-.425.288-.713T16 16q.425 0 .713.288T17 17q0 .425-.288.713T16 18Zm-7-6h6v-2H9v2ZM5 8h14V7H5v1Zm5-3h4V4h-4v1Z"/>`),
		g.Group(children),
	)
}