package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm-2 16q-.425 0-.713-.288T9 21V9H4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h16q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-5v12q0 .425-.288.713T14 22q-.425 0-.713-.288T13 21v-5h-2v5q0 .425-.288.713T10 22Z"/>`),
		g.Group(children),
	)
}