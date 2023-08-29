package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ebay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m415 92l-36 77l-33-77h-51l9 17q-2 0-7.5-.5t-8.5-.5q-21 0-33 9q-14 8-14 31h35q0-20 14-20t14 19v10q-40 0-52 10q-10-8-23-8q-20 0-29 14h-1v-67h-35v61q-7-4-19-6q-22-5-49-5q-58 0-76 15T2 214q0 30 18 45t76 15q27 0 49-5q6-1 19-6v26h34v-14q9 17 30 17q34 0 36-50h8q23 0 33-19h1l2 17h33q-2-16-2-24v-34l15 29v64h46v-64l62-119h-47zM96 176q14 0 21.5 5t8 9t.5 12H67q0-7 .5-11t7.5-9.5t21-5.5zm49 57h-21q0 21-28 21q-29 0-29-34h97v13h-19zm84-9q0 26-3 36t-12 10t-12-10t-3-36q0-12 .5-18t1.5-13t4.5-9.5t8.5-2.5t8.5 2.5t4.5 9.5t1.5 13t.5 18zm58-4q-14 0-14-20q0-6 1.5-10t2.5-6t5.5-3.5l6-2l8-1l7.5-.5v22.5l-1.5 8l-3 6.5l-5 4.5l-7.5 1.5z"/>`),
		g.Group(children),
	)
}