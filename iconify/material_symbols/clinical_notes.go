package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClinicalNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16q-1.25 0-2.125-.875T14 13q0-1.25.875-2.125T17 10q1.25 0 2.125.875T20 13q0 1.25-.875 2.125T17 16Zm-6 7v-2.9q0-.525.25-.988t.7-.737q.8-.475 1.688-.788t1.812-.462L17 19l1.55-1.875q.925.15 1.8.463t1.675.787q.45.275.713.738T23 20.1V23H11Zm-2-2.9v.9H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v5q-.775-.975-1.75-1.488T17 8V7H7v2h7q-.5.4-.9.9t-.675 1.1H7v2h5q0 .525.113 1.025t.312.975H7v2h3.45q-.675.575-1.063 1.388T9 20.1Z"/>`),
		g.Group(children),
	)
}