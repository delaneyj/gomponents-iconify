package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q.425 0 .713-.288T8 16q0-.425-.288-.713T7 15q-.425 0-.713.288T6 16q0 .425.288.713T7 17Zm-1-4h2V7H6v6Zm5 4h7v-2h-7v2Zm0-4h7v-2h-7v2Zm0-4h7V7h-7v2ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Z"/>`),
		g.Group(children),
	)
}