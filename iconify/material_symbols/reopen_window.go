package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReopenWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v7h-2V8H4v10h8v2H4Zm15 4q-1.825 0-3.188-1.137T14.1 20h1.55q.325 1.1 1.238 1.8t2.112.7q1.45 0 2.475-1.025T22.5 19q0-1.45-1.025-2.475T19 15.5q-.725 0-1.35.262t-1.1.738H18V18h-4v-4h1.5v1.425q.675-.65 1.575-1.038T19 14q2.075 0 3.538 1.463T24 19q0 2.075-1.463 3.538T19 24Z"/>`),
		g.Group(children),
	)
}