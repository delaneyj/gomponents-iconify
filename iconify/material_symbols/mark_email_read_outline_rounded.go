package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkEmailReadOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 6l8 5l8-5H4Zm0 14q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v4.35l-1 1l-1 1V8l-7.475 4.675q-.075.05-.525.15q-.125 0-.263-.037t-.262-.113L4 8v10h5.15l2 2H4Zm8-6Zm0-3Zm0 1.875Zm3.95 6.325l4.95-4.95q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-5.65 5.65q-.3.3-.7.3t-.7-.3l-2.85-2.85q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.15 2.15Z"/>`),
		g.Group(children),
	)
}