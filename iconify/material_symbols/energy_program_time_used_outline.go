package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergyProgramTimeUsedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 13Zm-7 9q-.825 0-1.413-.588T1 20V6q0-.825.588-1.413T3 4h8v2H3v14h14v-6h2v6q0 .825-.588 1.413T17 22H3Zm2-4h2v-7H5v7Zm4 0h2V8H9v10Zm4 0h2v-4h-2v4Zm5-6q-.75 0-1.475-.225t-1.35-.65l-.375.35q-.3.275-.713.275t-.687-.275q-.275-.275-.275-.7t.275-.7l.4-.4q-.4-.6-.6-1.275T13 7q0-2.075 1.463-3.537T18 2h5v5q0 2.075-1.463 3.538T18 12Zm0-2q1.25 0 2.125-.875T21 7V4h-3q-1.25 0-2.125.875T15 7q0 .325.063.625t.187.6l2.6-2.6q.275-.275.7-.275t.7.275q.3.3.3.713t-.3.712l-2.6 2.6q.325.15.663.25T18 10Zm-.15-2.975Z"/>`),
		g.Group(children),
	)
}