package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h2V9H6v6Zm5 0h2V9h-2v6Zm5 0h2V9h-2v6ZM4 17h16V7H4v10Zm0 0V7v10Zm1 4v-2H4q-.825 0-1.413-.588T2 17V7q0-.825.588-1.413T4 5h1V3h2v2h4V3h2v2h4V3h2v2h1q.825 0 1.413.588T22 7v10q0 .825-.588 1.413T20 19h-1v2h-2v-2h-4v2h-2v-2H7v2H5Z"/>`),
		g.Group(children),
	)
}