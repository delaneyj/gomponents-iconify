package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-.825 0-1.413-.588T7 19V9q0-.825.588-1.413T9 7h10q.825 0 1.413.588T21 9v10q0 .825-.588 1.413T19 21H9Zm0-10h10V9H9v2Zm4 4h2v-2h-2v2Zm0 4h2v-2h-2v2Zm-4-4h2v-2H9v2Zm8 0h2v-2h-2v2Zm-8 4h2v-2H9v2Zm8 0h2v-2h-2v2ZM5 17q-.825 0-1.413-.588T3 15V5q0-.825.588-1.413T5 3h10q.825 0 1.413.588T17 5v1h-2V5H5v10h1v2H5Z"/>`),
		g.Group(children),
	)
}