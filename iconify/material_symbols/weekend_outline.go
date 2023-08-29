package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeekendOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 7q0-1.25.875-2.125T7 4h10q1.25 0 2.125.875T20 7v2q1.25 0 2.125.875T23 12v5q0 1.25-.875 2.125T20 20H4q-1.25 0-2.125-.875T1 17v-5q0-1.25.875-2.125T4 9V7Zm2 0v2.8q.45.425.725.975T7 12v2h10v-2q0-.675.275-1.225T18 9.8V7q0-.425-.288-.713T17 6H7q-.425 0-.713.288T6 7Zm13 9H5v-4q0-.2-.075-.388t-.213-.324q-.137-.138-.325-.213T4 11q-.425 0-.713.288T3 12v5q0 .425.288.713T4 18h16q.425 0 .713-.288T21 17v-5q0-.425-.288-.713T20 11q-.2 0-.388.075t-.324.213q-.138.137-.213.324T19 12v4Zm-7 0Zm0 2Zm0-4Z"/>`),
		g.Group(children),
	)
}