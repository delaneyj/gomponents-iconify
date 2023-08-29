package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosswordOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16v4h4v-4h-4Zm-2-2v-4H4v4h4Zm2 0h4v-4h-4v4Zm6 0h4v-4h-4v4Zm0-6h4V4h-4v4Zm-8 8H4q-.825 0-1.413-.588T2 14v-4q0-.825.588-1.413T4 8h10V4q0-.825.588-1.413T16 2h4q.825 0 1.413.588T22 4v10q0 .825-.588 1.413T20 16h-4v4q0 .825-.588 1.413T14 22h-4q-.825 0-1.413-.588T8 20v-4Z"/>`),
		g.Group(children),
	)
}