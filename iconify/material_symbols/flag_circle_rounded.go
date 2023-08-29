package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCircleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 13H12l.725 1.45q.125.275.375.413t.525.137H17q.425 0 .713-.288T18 14v-4q0-.425-.288-.713T17 9h-2l-.725-1.45q-.125-.275-.375-.413T13.375 7H9q-.425 0-.713.288T8 8v9.25q0 .325.213.537T8.75 18q.325 0 .537-.213t.213-.537V13Zm3.5-2Zm-1 11q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}