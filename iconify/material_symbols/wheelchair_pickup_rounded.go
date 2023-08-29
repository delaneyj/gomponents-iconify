package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WheelchairPickupRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 6q-.825 0-1.413-.588T4.5 4q0-.825.588-1.413T6.5 2q.825 0 1.413.588T8.5 4q0 .825-.588 1.413T6.5 6ZM6 22q-.425 0-.713-.288T5 21v-6H4q-.425 0-.713-.288T3 14V9q0-.825.588-1.413T5 7h3q.825 0 1.413.588T10 9v1.95q-1.575.9-2.538 2.5T6.5 17q0 1.425.525 2.688T8.5 21.9v.1H6Zm7.5 0q-2.075 0-3.538-1.463T8.5 17q0-1.675.988-2.963T12 12.25v2.175q-.675.4-1.088 1.075T10.5 17q0 1.25.875 2.125T13.5 20q1.25 0 2.125-.875T16.5 17h2q0 2.075-1.463 3.538T13.5 22Zm6.875-3.125L18.45 16H14q-.425 0-.713-.288T13 15V9q0-.425.288-.713T14 8q.425 0 .713.288T15 9v5h4.025q.25 0 .475.125t.35.325l2.2 3.325q.225.35.138.75t-.438.625q-.35.225-.75.15t-.625-.425Z"/>`),
		g.Group(children),
	)
}