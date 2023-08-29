package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1Zm-6 5L0 9q2.375-2.425 5.488-3.713T12 4q3.4 0 6.513 1.288T24 9l-2.525 2.525q-.575-.275-1.2-.4T19 11q-2.5 0-4.25 1.75T13 17q0 .65.125 1.275t.4 1.2L12 21Zm5 0q-.425 0-.713-.288T16 20v-3q0-.425.288-.713T17 16v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1q.425 0 .713.288T22 17v3q0 .425-.288.713T21 21h-4Z"/>`),
		g.Group(children),
	)
}