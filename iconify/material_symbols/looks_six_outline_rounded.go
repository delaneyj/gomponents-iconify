package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksSixOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17h2q.825 0 1.413-.588T15 15v-2q0-.825-.588-1.413T13 11h-2V9h2q.425 0 .713-.288T14 8q0-.425-.288-.713T13 7h-2q-.825 0-1.413.588T9 9v6q0 .825.588 1.413T11 17Zm0-4h2v2h-2v-2Zm-6 8q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}