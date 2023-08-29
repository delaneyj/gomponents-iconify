package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupsThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 13.525L6.525 11L4 8.475L1.475 11L4 13.525ZM17.5 13L20 9l2.5 4h-5ZM1 18q-.425 0-.713-.288T0 17v-.575q0-1.1 1.1-1.763T4 14q.325 0 .613.025t.562.075q-.35.5-.513 1.075T4.5 16.4V18H1Zm6 0q-.425 0-.713-.288T6 17v-.6q0-1.625 1.663-2.638T12 12.75q2.7 0 4.35 1.012T18 16.4v.6q0 .425-.288.713T17 18H7Zm12.5 0v-1.6q0-.65-.175-1.225t-.5-1.075q.275-.05.563-.075T20 14q1.8 0 2.9.663t1.1 1.762V17q0 .425-.288.713T23 18h-3.5ZM12 12q-1.25 0-2.125-.875T9 9q0-1.25.875-2.125T12 6q1.25 0 2.125.875T15 9q0 1.25-.875 2.125T12 12Z"/>`),
		g.Group(children),
	)
}