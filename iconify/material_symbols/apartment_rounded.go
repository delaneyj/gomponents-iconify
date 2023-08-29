package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApartmentRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V9q0-.825.588-1.413T5 7h2V5q0-.825.588-1.413T9 3h6q.825 0 1.413.588T17 5v6h2q.825 0 1.413.588T21 13v6q0 .825-.588 1.413T19 21h-6v-4h-2v4H5Zm0-2h2v-2H5v2Zm0-4h2v-2H5v2Zm0-4h2V9H5v2Zm4 4h2v-2H9v2Zm0-4h2V9H9v2Zm0-4h2V5H9v2Zm4 8h2v-2h-2v2Zm0-4h2V9h-2v2Zm0-4h2V5h-2v2Zm4 12h2v-2h-2v2Zm0-4h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}