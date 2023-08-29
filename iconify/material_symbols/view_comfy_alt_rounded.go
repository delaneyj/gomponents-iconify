package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfyAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10H4q-.825 0-1.413-.588T2 8V4q0-.825.588-1.413T4 2h4q.825 0 1.413.588T10 4v4q0 .825-.588 1.413T8 10Zm0 12H4q-.825 0-1.413-.588T2 20v-4q0-.825.588-1.413T4 14h4q.825 0 1.413.588T10 16v4q0 .825-.588 1.413T8 22Zm12-12h-4q-.825 0-1.413-.588T14 8V4q0-.825.588-1.413T16 2h4q.825 0 1.413.588T22 4v4q0 .825-.588 1.413T20 10Zm0 12h-4q-.825 0-1.413-.588T14 20v-4q0-.825.588-1.413T16 14h4q.825 0 1.413.588T22 16v4q0 .825-.588 1.413T20 22Z"/>`),
		g.Group(children),
	)
}