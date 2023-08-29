package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-8h10v10H5Zm10 0v-2h2v2h-2Zm4 0v-2h2q0 .825-.588 1.413T19 21ZM4 19h8l-2.6-3.5L7.5 18l-1.4-1.85L4 19Zm15-2v-2h2v2h-2Zm0-4v-2h2v2h-2ZM3 9V7h2v2H3Zm16 0V7h2v2h-2ZM3 5q0-.825.588-1.413T5 3v2H3Zm4 0V3h2v2H7Zm4 0V3h2v2h-2Zm4 0V3h2v2h-2Zm4 0V3q.825 0 1.413.588T21 5h-2Z"/>`),
		g.Group(children),
	)
}