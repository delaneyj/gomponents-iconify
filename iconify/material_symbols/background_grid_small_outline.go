package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackgroundGridSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h2v-2H5v2Zm4 0h2v-2H9v2Zm4 0h2v-2h-2v2Zm4 0h2v-2h-2v2ZM5 7h2V5H5v2Zm0 4h2V9H5v2Zm0 4h2v-2H5v2Zm4-8h2V5H9v2Zm0 4h2V9H9v2Zm0 4h2v-2H9v2Zm4-8h2V5h-2v2Zm0 4h2V9h-2v2Zm0 4h2v-2h-2v2Zm4-8h2V5h-2v2Zm0 4h2V9h-2v2Zm0 4h2v-2h-2v2ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}