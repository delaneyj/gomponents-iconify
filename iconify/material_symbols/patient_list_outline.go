package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientListOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 14q-1.25 0-2.125-.875T13 11q0-1.25.875-2.125T16 8q1.25 0 2.125.875T19 11q0 1.25-.875 2.125T16 14Zm-6 6v-1.9q0-.525.25-1t.7-.75q1.125-.675 2.388-1.012T16 15q1.4 0 2.663.338t2.387 1.012q.45.275.7.75t.25 1V20H10Zm2.15-2h7.7q-.875-.5-1.85-.75T16 17q-1.025 0-2 .25t-1.85.75ZM16 12q.425 0 .713-.288T17 11q0-.425-.288-.713T16 10q-.425 0-.713.288T15 11q0 .425.288.713T16 12Zm0-1Zm0 7ZM3 14v-2h8v2H3Zm0-8V4h12v2H3Zm8.1 4H3V8h9q-.35.425-.563.925T11.1 10Z"/>`),
		g.Group(children),
	)
}