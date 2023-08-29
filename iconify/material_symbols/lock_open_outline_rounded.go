package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V10q0-.825.588-1.413T6 8h9V6q0-1.25-.875-2.125T12 3q-1.05 0-1.838.638T9.1 5.224q-.1.35-.413.563T8 6q-.425 0-.713-.275t-.212-.65q.35-1.725 1.725-2.9T12 1q2.075 0 3.538 1.462T17 6v2h1q.825 0 1.413.588T20 10v10q0 .825-.588 1.413T18 22H6Zm0-2h12V10H6v10Zm6-3q.825 0 1.413-.588T14 15q0-.825-.588-1.413T12 13q-.825 0-1.413.588T10 15q0 .825.588 1.413T12 17Zm-6 3V10v10Z"/>`),
		g.Group(children),
	)
}