package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfflineShareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23q-.825 0-1.413-.588T4 21V6q0-.425.288-.713T5 5q.425 0 .713.288T6 6v15h9q.425 0 .713.288T16 22q0 .425-.288.713T15 23H6Zm4-4q-.825 0-1.413-.588T8 17V3q0-.825.588-1.413T10 1h8q.825 0 1.413.588T20 3v14q0 .825-.588 1.413T18 19h-8Zm0-5h8V6h-8v8Zm4.15-3.75H12.5v1q0 .325-.213.537T11.75 12q-.325 0-.537-.213T11 11.25v-1.5q0-.425.288-.713T12 8.75h2.15l-.175-.175q-.225-.225-.225-.525t.225-.525Q14.2 7.3 14.5 7.3t.525.225L16.65 9.15q.15.15.15.35t-.15.35l-1.625 1.625q-.225.225-.525.225t-.525-.225q-.225-.225-.225-.525t.225-.525l.175-.175Z"/>`),
		g.Group(children),
	)
}