package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergyProgramTimeUsed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.45 8.425q.3.3.7.313t.7-.288l1.4-1.4q.3-.3.3-.712t-.3-.713q-.275-.275-.7-.275t-.7.275l-1.4 1.4q-.275.275-.275.688t.275.712ZM3 22q-.825 0-1.413-.588T1 20V6q0-.825.588-1.413T3 4h8v8q0 .825.588 1.413T13 14h6v6q0 .825-.588 1.413T17 22H3Zm15-10q-.75 0-1.475-.225t-1.35-.65l-.375.35q-.3.275-.713.275t-.687-.275q-.275-.275-.275-.7t.275-.7l.4-.4q-.4-.6-.6-1.275T13 7q0-2.075 1.463-3.537T18 2h5v5q0 2.075-1.463 3.538T18 12ZM5 18h2v-7H5v7Zm4 0h2V8H9v10Zm4 0h2v-4h-2v4Z"/>`),
		g.Group(children),
	)
}