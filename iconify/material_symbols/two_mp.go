package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoMp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 18.5h1.5V14h1v3H10v-3h1v4.5h1.5v-5q0-.425-.288-.713T11.5 12.5H7q-.425 0-.713.288T6 13.5v5Zm3.75-7h4.5V10h-3V9h2q.425 0 .713-.288T14.25 8V6.5q0-.425-.288-.713T13.25 5.5h-3.5V7h3v1h-2q-.425 0-.713.288T9.75 9v2.5Zm3.75 7H15V17h2q.425 0 .713-.288T18 16v-2.5q0-.425-.288-.713T17 12.5h-3.5v6Zm1.5-3V14h1.5v1.5H15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}