package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArOnYouOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4q3.35 0 5.675 2.325T20 12q0 3.35-2.325 5.675T12 20q-3.35 0-5.675-2.325T4 12q0-3.35 2.325-5.675T12 4Zm0 14q2.5 0 4.25-1.75T18 12q0-2.5-1.75-4.25T12 6Q9.5 6 7.75 7.75T6 12q0 2.5 1.75 4.25T12 18Zm0-2.5q1.2 0 2.15-.688T15.5 13h-7q.4 1.125 1.35 1.813T12 15.5ZM8.5 10q0 .425.288.713T9.5 11q.425 0 .713-.288T10.5 10q0-.425-.288-.713T9.5 9q-.425 0-.713.288T8.5 10Zm5 0q0 .425.288.713T14.5 11q.425 0 .713-.288T15.5 10q0-.425-.288-.713T14.5 9q-.425 0-.713.288T13.5 10ZM1 6V1h5v2H3v3H1Zm0 17v-5h2v3h3v2H1Zm17 0v-2h3v-3h2v5h-5Zm3-17V3h-3V1h5v5h-2Zm-9 6Z"/>`),
		g.Group(children),
	)
}