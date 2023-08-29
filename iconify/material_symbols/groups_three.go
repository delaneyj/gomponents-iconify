package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupsThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13.525L6.525 11L4 8.475L1.475 11L4 13.525ZM17.5 13L20 9l2.5 4h-5ZM12 12q-1.25 0-2.125-.875T9 9q0-1.275.875-2.138T12 6q1.275 0 2.138.863T15 9q0 1.25-.863 2.125T12 12ZM0 18v-1.575q0-1.1 1.113-1.763T4 14q.325 0 .625.013t.575.062q-.35.5-.525 1.075T4.5 16.375V18H0Zm6 0v-1.625q0-1.625 1.663-2.625t4.337-1q2.7 0 4.35 1T18 16.375V18H6Zm14-4q1.8 0 2.9.663t1.1 1.762V18h-4.5v-1.625q0-.65-.163-1.225t-.487-1.075q.275-.05.563-.062T20 14Z"/>`),
		g.Group(children),
	)
}