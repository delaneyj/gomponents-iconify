package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinearScale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17q-1.8 0-3.175-1.137T12.1 13H6.8q-.3.675-.925 1.088T4.5 14.5q-1.05 0-1.775-.725T2 12q0-1.05.725-1.775T4.5 9.5q.75 0 1.375.413T6.8 11h5.3q.35-1.725 1.725-2.863T17 7q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17Zm0-2q1.25 0 2.125-.875T20 12q0-1.25-.875-2.125T17 9q-1.25 0-2.125.875T14 12q0 1.25.875 2.125T17 15Z"/>`),
		g.Group(children),
	)
}