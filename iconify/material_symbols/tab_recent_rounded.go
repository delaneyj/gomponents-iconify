package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabRecentRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 17.8v-2.3q0-.2-.15-.35T18 15q-.2 0-.35.15t-.15.35v2.275q0 .2.075.388t.225.337l1.5 1.5q.15.15.35.15T20 20q.15-.15.15-.35T20 19.3l-1.5-1.5ZM13 6v3q0 .425.288.713T14 10h6V6h-7Zm5 17q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6.275q-.875-.625-1.913-.95T17.975 11q-2.9 0-4.938 2.05T11 18q0 .525.075 1.025T11.3 20H4Z"/>`),
		g.Group(children),
	)
}