package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.825 0-1.413-.588T3 18q0-.825.588-1.413T5 16h14q.825 0 1.413.588T21 18q0 .825-.588 1.413T19 20H5ZM5 8q-.825 0-1.413-.588T3 6q0-.825.588-1.413T5 4h14q.825 0 1.413.588T21 6q0 .825-.588 1.413T19 8H5Zm0 6q-.825 0-1.413-.588T3 12q0-.825.588-1.413T5 10h14q.825 0 1.413.588T21 12q0 .825-.588 1.413T19 14H5Zm1-7q.425 0 .713-.288T7 6q0-.425-.288-.713T6 5q-.425 0-.713.288T5 6q0 .425.288.713T6 7Zm0 6q.425 0 .713-.288T7 12q0-.425-.288-.713T6 11q-.425 0-.713.288T5 12q0 .425.288.713T6 13Zm0 6q.425 0 .713-.288T7 18q0-.425-.288-.713T6 17q-.425 0-.713.288T5 18q0 .425.288.713T6 19Z"/>`),
		g.Group(children),
	)
}