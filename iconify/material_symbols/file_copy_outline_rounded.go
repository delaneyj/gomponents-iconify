package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCopyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19H8q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h6.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V17q0 .825-.588 1.413T19 19Zm0-11h-3.5q-.625 0-1.063-.438T14 6.5V3H8v14h11V8ZM4 23q-.825 0-1.413-.588T2 21V8q0-.425.288-.713T3 7q.425 0 .713.288T4 8v13h10q.425 0 .713.288T15 22q0 .425-.288.713T14 23H4ZM8 3v5v-5v14V3Z"/>`),
		g.Group(children),
	)
}