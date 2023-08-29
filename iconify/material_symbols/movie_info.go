package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Zm0-2h2v-2H4v2Zm14 0h2v-2h-2v2Zm-7-2h2v-6h-2v6Zm-7-2h2v-2H4v2Zm14 0h2v-2h-2v2ZM4 11h2V9H4v2Zm14 0h2V9h-2v2Zm-6-2q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9ZM4 7h2V5H4v2Zm14 0h2V5h-2v2Z"/>`),
		g.Group(children),
	)
}