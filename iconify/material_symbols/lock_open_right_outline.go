package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20h12V10H6v10Zm6-3q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13q-.825 0-1.413.588T10 15q0 .825.588 1.413T12 17Zm-6 3V10v10Zm0 2q-.825 0-1.413-.588T4 20V10q0-.825.588-1.413T6 8h7V6q0-2.075 1.463-3.538T18 1q2.075 0 3.538 1.463T23 6h-2q0-1.25-.875-2.125T18 3q-1.25 0-2.125.875T15 6v2h3q.825 0 1.413.588T20 10v10q0 .825-.588 1.413T18 22H6Z"/>`),
		g.Group(children),
	)
}