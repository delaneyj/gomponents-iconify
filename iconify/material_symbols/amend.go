package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 16h10v-2H7v2Zm0-4h4v-1.5H9.575q.5-.475 1.125-.738T12 9.5q1.175 0 2.087.688T15.35 12h1.55q-.35-1.75-1.725-2.875T12 8q-.975 0-1.887.375T8.5 9.45V8H7v4Zm5 10q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}