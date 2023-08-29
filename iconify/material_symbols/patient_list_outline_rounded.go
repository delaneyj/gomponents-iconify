package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientListOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 14q-1.25 0-2.125-.875T13 11q0-1.25.875-2.125T16 8q1.25 0 2.125.875T19 11q0 1.25-.875 2.125T16 14Zm-5 6q-.425 0-.713-.288T10 19v-.9q0-.525.25-1t.7-.75q1.125-.675 2.388-1.012T16 15q1.4 0 2.663.338t2.387 1.012q.45.275.7.75t.25 1v.9q0 .425-.288.713T21 20H11Zm1.15-2h7.7q-.875-.5-1.85-.75T16 17q-1.025 0-2 .25t-1.85.75ZM16 12q.425 0 .713-.288T17 11q0-.425-.288-.713T16 10q-.425 0-.713.288T15 11q0 .425.288.713T16 12Zm0-1Zm0 7Zm-6-4H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h6q.425 0 .713.288T11 13q0 .425-.288.713T10 14Zm4-8H4q-.425 0-.713-.288T3 5q0-.425.288-.713T4 4h10q.425 0 .713.288T15 5q0 .425-.288.713T14 6Zm-2.9 4H4q-.425 0-.713-.288T3 9q0-.425.288-.713T4 8h8q-.35.425-.563.925T11.1 10Z"/>`),
		g.Group(children),
	)
}