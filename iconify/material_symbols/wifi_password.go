package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiPassword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 24q-.425 0-.713-.288T18 23v-3q0-.425.288-.713T19 19v-1q0-.825.588-1.413T21 16q.825 0 1.413.588T23 18v1q.425 0 .713.288T24 20v3q0 .425-.288.713T23 24h-4Zm1-5h2v-1q0-.425-.288-.713T21 17q-.425 0-.713.288T20 18v1ZM6.35 15.35l-2.1-2.15q1.475-1.475 3.463-2.337T12 10q2.3 0 4.288.875t3.462 2.375l-2.1 2.1q-1.1-1.1-2.55-1.725T12 13q-1.65 0-3.1.625T6.35 15.35ZM2.1 11.1L0 9q2.3-2.35 5.375-3.675T12 4q3.55 0 6.625 1.325T24 9l-2.1 2.1q-1.925-1.925-4.463-3.013T12 7Q9.1 7 6.562 8.088T2.1 11.1ZM12 21l3.525-3.55q-.675-.675-1.575-1.063T12 16q-1.05 0-1.95.388T8.475 17.45L12 21Z"/>`),
		g.Group(children),
	)
}