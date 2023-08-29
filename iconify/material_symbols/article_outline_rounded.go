package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h5q.425 0 .713-.288T14 16q0-.425-.288-.713T13 15H8q-.425 0-.713.288T7 16q0 .425.288.713T8 17Zm0-4h8q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13Zm0-4h8q.425 0 .713-.288T17 8q0-.425-.288-.713T16 7H8q-.425 0-.713.288T7 8q0 .425.288.713T8 9ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}