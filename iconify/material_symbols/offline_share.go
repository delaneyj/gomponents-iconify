package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfflineShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23q-.825 0-1.413-.588T4 21V5h2v16h10v2H6Zm4-4q-.825 0-1.413-.588T8 17V3q0-.825.588-1.413T10 1h8q.825 0 1.413.588T20 3v14q0 .825-.588 1.413T18 19h-8Zm0-5h8V6h-8v8Zm1-2V9.75q0-.425.288-.713T12 8.75h2.15l-.7-.7L14.5 7L17 9.5L14.5 12l-1.05-1.05l.7-.7H12.5V12H11Z"/>`),
		g.Group(children),
	)
}