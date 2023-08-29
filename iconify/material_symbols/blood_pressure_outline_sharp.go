package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodPressureOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.5h-2V6H4v3H2Zm8.675 11H4q-.825 0-1.413-.588T2 18v-3h2v3h6.075q.075.525.225 1.025t.375.975ZM12 12Zm5 10q-2.075 0-3.538-1.463T12 17q0-2.075 1.463-3.538T17 12q2.075 0 3.538 1.463T22 17q0 2.075-1.463 3.538T17 22Zm.2-4.5l2.275-2.275l-.7-.7L16.5 16.8l.7.7ZM2 13v-2h3.6L7 13.775L10.35 7h1.275l1.95 3.9q-.45.275-.863.575t-.762.675l-.95-1.9L7.625 17h-1.25l-2-4H2Z"/>`),
		g.Group(children),
	)
}