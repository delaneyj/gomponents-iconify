package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiTwoBarLockedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.375-2.425 5.488-3.713T12 4q3.425 0 6.525 1.275T24 9l-2.525 2.525q-.55-.25-1.137-.375T19.15 11l1.95-1.95q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l4.3 4.3q1.1-.65 2.312-1T12 12q.775 0 1.538.163t1.487.362q-.95.875-1.488 2.038T13 17.024q0 .65.125 1.263t.4 1.187L12 21Zm4 0v-5h1v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1Z"/>`),
		g.Group(children),
	)
}