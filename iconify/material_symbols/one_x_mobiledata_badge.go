package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneXMobiledataBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm4-4h2V7H5v2h2v8Zm4.5 0h2l1.75-3.175L17 17h2l-2.75-5L19 7h-2l-1.75 3.175L13.5 7h-2l2.75 5l-2.75 5Z"/>`),
		g.Group(children),
	)
}