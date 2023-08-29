package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowsheetSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11h5V9H6v2Zm-4 9V4h19.9v1.375q-.425-.2-.912-.288T20 5q-2.075 0-3.538 1.463T15 10q0 .55.1 1.05t.3.95H11v1H6v2h5v1q-.95.675-1.475 1.713T9 20H2Zm12 1q.425 0 .713-.288T15 20q0-.425-.288-.713T14 19q-.425 0-.713.288T13 20q0 .425.288.713T14 21Zm6-10q.425 0 .713-.288T21 10q0-.425-.288-.713T20 9q-.425 0-.713.288T19 10q0 .425.288.713T20 11Zm-6 12q-1.25 0-2.125-.875T11 20q0-.975.563-1.75T13 17.175V14h6v-1.175q-.875-.3-1.438-1.075T17 10q0-1.25.875-2.125T20 7q1.25 0 2.125.875T23 10q0 .975-.563 1.75T21 12.825V16h-6v1.175q.875.3 1.438 1.075T17 20q0 1.25-.875 2.125T14 23Z"/>`),
		g.Group(children),
	)
}