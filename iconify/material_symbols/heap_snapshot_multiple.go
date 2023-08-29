package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapSnapshotMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 11.5q.525 0 .888-.363t.362-.887q0-.525-.363-.888T11.25 9q-.525 0-.888.363T10 10.25q0 .525.363.888t.887.362Zm.175 4.5L17 10.425L15.575 9L10 14.575L11.425 16Zm4.325 0q.525 0 .888-.363T17 14.75q0-.525-.363-.888t-.887-.362q-.525 0-.888.363t-.362.887q0 .525.363.888t.887.362ZM19 19H8q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h7l6 6v10q0 .825-.588 1.413T19 19ZM14 8h5l-5-5v5ZM4 23q-.825 0-1.413-.588T2 21V7h2v14h11v2H4Z"/>`),
		g.Group(children),
	)
}