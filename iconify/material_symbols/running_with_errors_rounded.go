package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWithErrorsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q1.875 0 3.688.7t3.162 2.05q.3.3.3.7t-.3.7L13.7 11.3q-.475.475-1.087.213T12 10.575V4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20q1.725 0 3.3-.713T18 17.25V20q-1.325.95-2.85 1.475T12 22Zm9-4q-.425 0-.713-.288T20 17v-6q0-.425.288-.713T21 10q.425 0 .713.288T22 11v6q0 .425-.288.713T21 18Zm0 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Z"/>`),
		g.Group(children),
	)
}