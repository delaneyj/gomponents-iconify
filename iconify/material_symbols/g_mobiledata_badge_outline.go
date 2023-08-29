package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GMobiledataBadgeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Zm5 12h4q.825 0 1.413-.588T16 15v-4h-4v2h2v2h-4V9h6q0-.825-.588-1.413T14 7h-4q-.825 0-1.413.588T8 9v6q0 .825.588 1.413T10 17Z"/>`),
		g.Group(children),
	)
}