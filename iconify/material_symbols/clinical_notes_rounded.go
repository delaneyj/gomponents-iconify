package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalNotesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q-1.25 0-2.125-.875T14 13q0-1.25.875-2.125T17 10q1.25 0 2.125.875T20 13q0 1.25-.875 2.125T17 16Zm-8 4.1v.9H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v5q-.775-.975-1.75-1.488T17 8q0-.425-.288-.713T16 7H8q-.425 0-.713.288T7 8q0 .425.288.713T8 9h6q-.5.4-.9.9t-.675 1.1H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13h4q0 .525.113 1.025t.312.975H8q-.425 0-.713.288T7 16q0 .425.288.713T8 17h2.45q-.675.575-1.063 1.388T9 20.1Zm3 2.9q-.425 0-.713-.288T11 22v-1.9q0-.525.25-.988t.7-.737q.8-.475 1.35-.663t1.475-.337q.3-.05.6.013t.5.287L17 19l1.1-1.325q.2-.25.5-.3t.6 0q.925.15 1.475.338t1.35.662q.45.275.713.738T23 20.1V22q0 .425-.288.713T22 23H12Z"/>`),
		g.Group(children),
	)
}