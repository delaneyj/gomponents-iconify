package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArOnYou(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4q3.35 0 5.675 2.325T20 12q0 3.35-2.325 5.675T12 20q-3.35 0-5.675-2.325T4 12q0-3.35 2.325-5.675T12 4Zm0 11.5q1.2 0 2.15-.688T15.5 13h-7q.4 1.125 1.35 1.813T12 15.5ZM8.5 10q0 .425.288.713T9.5 11q.425 0 .713-.288T10.5 10q0-.425-.288-.713T9.5 9q-.425 0-.713.288T8.5 10Zm5 0q0 .425.288.713T14.5 11q.425 0 .713-.288T15.5 10q0-.425-.288-.713T14.5 9q-.425 0-.713.288T13.5 10ZM1 6V4q0-1.25.875-2.125T4 1h2v2H4q-.425 0-.713.288T3 4v2H1Zm3 17q-1.25 0-2.125-.875T1 20v-2h2v2q0 .425.288.713T4 21h2v2H4Zm14 0v-2h2q.425 0 .713-.288T21 20v-2h2v2q0 1.25-.875 2.125T20 23h-2Zm3-17V4q0-.425-.288-.713T20 3h-2V1h2q1.25 0 2.125.875T23 4v2h-2Z"/>`),
		g.Group(children),
	)
}