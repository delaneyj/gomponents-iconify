package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6ZM9 22V9q-1.5-.125-3.05-.375T3 8l.5-2q1.95.525 4.15.763T12 7q2.15 0 4.35-.238T20.5 6l.5 2q-1.4.375-2.95.625T15 9v13h-2v-6h-2v6H9Z"/>`),
		g.Group(children),
	)
}