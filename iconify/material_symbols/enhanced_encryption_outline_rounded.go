package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnhancedEncryptionOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16v2q0 .425.288.713T12 19q.425 0 .713-.288T13 18v-2h2q.425 0 .713-.288T16 15q0-.425-.288-.713T15 14h-2v-2q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12v2H9q-.425 0-.713.288T8 15q0 .425.288.713T9 16h2Zm-5 6q-.825 0-1.413-.588T4 20V10q0-.825.588-1.413T6 8h1V6q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6v2h1q.825 0 1.413.588T20 10v10q0 .825-.588 1.413T18 22H6Zm0-2h12V10H6v10ZM9 8h6V6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6v2ZM6 20V10v10Z"/>`),
		g.Group(children),
	)
}