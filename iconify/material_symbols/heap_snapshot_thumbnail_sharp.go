package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeapSnapshotThumbnailSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h14v14H1Zm16-8V5h6v6h-6Zm0 8v-6h6v6h-6Zm-6.5-3q.625 0 1.063-.438T12 14.5q0-.625-.438-1.063T10.5 13q-.625 0-1.063.438T9 14.5q0 .625.438 1.063T10.5 16Zm-5.05-.05l6.5-6.5l-1.4-1.4l-6.5 6.5l1.4 1.4ZM5.5 11q.625 0 1.062-.438T7 9.5q0-.625-.438-1.063T5.5 8q-.625 0-1.063.438T4 9.5q0 .625.438 1.063T5.5 11Z"/>`),
		g.Group(children),
	)
}