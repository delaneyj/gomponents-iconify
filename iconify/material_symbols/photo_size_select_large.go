package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V7h14v14H5Zm14 0v-2h2q0 .825-.588 1.413T19 21ZM5 18h10l-3.4-4.5L9 17l-1.6-2.15L5 18Zm14-1v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2ZM3 5q0-.825.588-1.413T5 3v2H3Zm4 0V3h2v2H7Zm4 0V3h2v2h-2Zm4 0V3h2v2h-2Zm4 0V3q.825 0 1.413.588T21 5h-2Z"/>`),
		g.Group(children),
	)
}