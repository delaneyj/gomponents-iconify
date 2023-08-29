package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapSnapshotMultipleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 11.5q.525 0 .888-.363t.362-.887q0-.525-.363-.888T11.25 9q-.525 0-.888.363T10 10.25q0 .525.363.888t.887.362Zm.175 4.5L17 10.425L15.575 9L10 14.575L11.425 16Zm4.325 0q.525 0 .888-.363T17 14.75q0-.525-.363-.888t-.887-.362q-.525 0-.888.363t-.362.887q0 .525.363.888t.887.362ZM6 19V1h9l6 6v12H6Zm8-11V3H8v14h11V8h-5ZM2 23V7h2v14h11v2H2ZM8 3v5v-5v14V3Z"/>`),
		g.Group(children),
	)
}