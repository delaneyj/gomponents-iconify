package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapSnapshotLargeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 18.5q.625 0 1.063-.438T16 17q0-.625-.438-1.063T14.5 15.5q-.625 0-1.063.438T13 17q0 .625.438 1.063t1.062.437Zm-5.05-.05l6.5-6.5l-1.4-1.4l-6.5 6.5l1.4 1.4Zm.05-4.95q.625 0 1.063-.438T11 12q0-.625-.438-1.063T9.5 10.5q-.625 0-1.063.438T8 12q0 .625.438 1.063T9.5 13.5ZM4 22V2h10l6 6v14H4Zm9-13V4H6v16h12V9h-5ZM6 4v5v-5v16V4Z"/>`),
		g.Group(children),
	)
}