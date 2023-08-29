package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WcRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 22q-.425 0-.713-.288T5.5 21v-6.5H5q-.425 0-.713-.288T4 13.5V9q0-.825.588-1.413T6 7h3q.825 0 1.413.588T11 9v4.5q0 .425-.288.713T10 14.5h-.5V21q0 .425-.288.713T8.5 22h-2Zm9.5 0q-.425 0-.713-.288T15 21v-5h-1.625q-.5 0-.8-.413t-.125-.912l2.1-6.325q.2-.65.737-1T16.5 7q.675 0 1.213.35t.737 1l2.1 6.325q.175.5-.125.913t-.8.412H18v5q0 .425-.288.713T17 22h-1ZM7.5 6q-.825 0-1.413-.588T5.5 4q0-.825.588-1.413T7.5 2q.825 0 1.413.588T9.5 4q0 .825-.588 1.413T7.5 6Zm9 0q-.825 0-1.413-.588T14.5 4q0-.825.588-1.413T16.5 2q.825 0 1.413.588T18.5 4q0 .825-.588 1.413T16.5 6Z"/>`),
		g.Group(children),
	)
}