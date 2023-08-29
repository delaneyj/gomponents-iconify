package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWithErrors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.2 0 4.163.9t3.387 2.55L12 13V4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20q1.725 0 3.3-.713T18 17.25V20q-1.325.95-2.85 1.475T12 22Zm8-4v-8h2v8h-2Zm1 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Z"/>`),
		g.Group(children),
	)
}