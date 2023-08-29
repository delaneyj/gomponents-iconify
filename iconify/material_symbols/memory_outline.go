package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 15V9h6v6H9Zm2-2h2v-2h-2v2Zm-2 8v-2H7q-.825 0-1.413-.588T5 17v-2H3v-2h2v-2H3V9h2V7q0-.825.588-1.413T7 5h2V3h2v2h2V3h2v2h2q.825 0 1.413.588T19 7v2h2v2h-2v2h2v2h-2v2q0 .825-.588 1.413T17 19h-2v2h-2v-2h-2v2H9Zm8-4V7H7v10h10Zm-5-5Z"/>`),
		g.Group(children),
	)
}