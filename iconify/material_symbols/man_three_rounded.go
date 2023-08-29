package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22.25q-.425 0-.713-.288T10 21.25v-6H9q-.425 0-.713-.288T8 14.25v-5q0-.825.588-1.413T10 7.25h4q.825 0 1.413.588T16 9.25v5q0 .425-.288.713T15 15.25h-1v6q0 .425-.288.713T13 22.25h-2Zm.65-16.1L10.1 4.6q-.15-.15-.15-.35t.15-.35l1.55-1.55q.15-.15.35-.15t.35.15L13.9 3.9q.15.15.15.35t-.15.35l-1.55 1.55q-.15.15-.35.15t-.35-.15Z"/>`),
		g.Group(children),
	)
}