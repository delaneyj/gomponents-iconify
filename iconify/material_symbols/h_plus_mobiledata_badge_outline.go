package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HPlusMobiledataBadgeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm0-2h18V5H3v14Zm0 0V5v14Zm1.5-2h2v-4h4v4h2V7h-2v4h-4V7h-2v10Zm11-2h2v-2h2v-2h-2V9h-2v2h-2v2h2v2Z"/>`),
		g.Group(children),
	)
}