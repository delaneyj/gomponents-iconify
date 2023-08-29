package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.75 14.95L16.3 13.5q.225-.375.388-.775t.237-.825q.075-.375.363-.638t.637-.262q.475 0 .775.325t.225.75q-.125.8-.425 1.513t-.75 1.362Zm-2.95-3L9 6.15V5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5v6q0 .275-.063.5t-.137.45Zm4.3 9.95l-17-17q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM11 20v-2.1q-2.325-.3-3.95-1.925t-1.975-3.9q-.075-.425.225-.75T6.1 11q.35 0 .625.263t.35.637q.325 1.75 1.7 2.925T12 16q.85 0 1.613-.263T15 15l1.425 1.425q-.725.575-1.588.963T13 17.9V20q0 .425-.288.713T12 21q-.425 0-.713-.288T11 20Z"/>`),
		g.Group(children),
	)
}