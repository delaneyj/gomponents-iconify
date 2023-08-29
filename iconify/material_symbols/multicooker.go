package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multicooker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 5V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v1h3q.825 0 1.413.588T21 7v1H3V7q0-.825.588-1.413T5 5h3ZM5 21q-.825 0-1.413-.588T3 19v-9h4v2q0 .825.588 1.413T9 14h6q.825 0 1.413-.588T17 12v-2h4v9q0 .825-.588 1.413T19 21H5Zm3-3q.425 0 .713-.288T9 17q0-.425-.288-.713T8 16q-.425 0-.713.288T7 17q0 .425.288.713T8 18Zm4 0q.425 0 .713-.288T13 17q0-.425-.288-.713T12 16q-.425 0-.713.288T11 17q0 .425.288.713T12 18Zm4 0q.425 0 .713-.288T17 17q0-.425-.288-.713T16 16q-.425 0-.713.288T15 17q0 .425.288.713T16 18Zm-7-6v-2h6v2H9Zm1-7h4V4h-4v1Z"/>`),
		g.Group(children),
	)
}