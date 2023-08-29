package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h1.5v-4.5h1v3H10v-3h1V15h1.5v-5q0-.425-.288-.713T11.5 9H7q-.425 0-.713.288T6 10v5Zm7.5 0H15v-1.5h2q.425 0 .713-.288T18 12.5V10q0-.425-.288-.713T17 9h-3.5v6Zm1.5-3v-1.5h1.5V12H15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}