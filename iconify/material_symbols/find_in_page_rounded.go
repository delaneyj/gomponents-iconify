package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FindInPageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8.075q.45 0 .85.188t.675.537l3.925 4.725q.225.275.35.6t.125.675V19.65L15.45 15q.275-.425.413-.925T16 13q0-1.65-1.175-2.825T12 9q-1.65 0-2.825 1.175T8 13q0 1.65 1.175 2.825T12 17q.575 0 1.075-.138T14 16.45l5.2 5.15q-.575.25-1.2.325T16.75 22H6Zm6-7q-.825 0-1.413-.588T10 13q0-.825.588-1.413T12 11q.825 0 1.413.588T14 13q0 .825-.588 1.413T12 15Z"/>`),
		g.Group(children),
	)
}