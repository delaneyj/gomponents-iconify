package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapSnapshotLargeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 18.5q.625 0 1.063-.438T16 17q0-.625-.438-1.063T14.5 15.5q-.625 0-1.063.438T13 17q0 .625.438 1.063t1.062.437Zm-5.75-.75q.275.275.7.275t.7-.275l5.1-5.1q.275-.275.275-.7t-.275-.7q-.275-.275-.7-.275t-.7.275l-5.1 5.1q-.275.275-.275.7t.275.7Zm.75-4.25q.625 0 1.063-.438T11 12q0-.625-.438-1.063T9.5 10.5q-.625 0-1.063.438T8 12q0 .625.438 1.063T9.5 13.5ZM6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v12q0 .825-.588 1.413T18 22H6Zm7-15V4H6v16h12V9h-3q-.825 0-1.413-.588T13 7ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}