package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewWeek(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h1.325q.825 0 1.413.588T7.325 6v12q0 .825-.588 1.413T5.325 20H4Zm7.35 0q-.825 0-1.412-.588T9.35 18V6q0-.825.588-1.413T11.35 4h1.325q.825 0 1.413.588T14.675 6v12q0 .825-.588 1.413T12.675 20H11.35Zm7.325 0q-.825 0-1.413-.588T16.675 18V6q0-.825.588-1.413T18.674 4H20q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20h-1.325Z"/>`),
		g.Group(children),
	)
}