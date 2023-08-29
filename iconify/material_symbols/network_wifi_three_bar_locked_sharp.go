package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiThreeBarLockedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L0 9q2.4-2.45 5.5-3.725T12 4q3.425 0 6.525 1.275T24 9l-2.525 2.525q-.55-.25-1.125-.375t-1.2-.15l1.95-1.95q-1.95-1.475-4.263-2.263T12 6q-2.525 0-4.838.788T2.9 9.05l2.9 2.9q1.325-.95 2.9-1.487t3.3-.538q1.4 0 2.7.363t2.45 1.012q-1.825.575-2.988 2.138T13 17.024q0 .65.125 1.263t.4 1.187L12 21Zm4 0v-5h1v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1Z"/>`),
		g.Group(children),
	)
}