package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberSmartRecordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19q-2.925 0-4.963-2.038T2 12q0-2.925 2.038-4.963T9 5q2.925 0 4.963 2.038T16 12q0 2.925-2.038 4.963T9 19Zm0-7Zm7 6.9v-2q1.75-.35 2.875-1.725T20 12q0-1.8-1.125-3.175T16 7.1v-2q2.6.35 4.3 2.312T22 12q0 2.625-1.7 4.588T16 18.9ZM9 17q2.075 0 3.538-1.463T14 12q0-2.075-1.463-3.538T9 7Q6.925 7 5.462 8.463T4 12q0 2.075 1.463 3.538T9 17Z"/>`),
		g.Group(children),
	)
}