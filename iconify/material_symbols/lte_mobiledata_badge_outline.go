package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LteMobiledataBadgeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm0-2h18V5H3v14Zm0 0V5v14Zm1-3h5v-2H6V8H4v8Zm6 0h2v-6h2V8H8v2h2v6Zm5 0h5v-2h-3v-1h2v-2h-2v-1h3V8h-5v8Z"/>`),
		g.Group(children),
	)
}