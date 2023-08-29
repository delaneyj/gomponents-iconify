package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewInAr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19.475L6 16.6q-.475-.275-.738-.738T5 14.876v-5.75q0-.525.263-.988T6 7.4l5-2.875q.475-.275 1-.275t1 .275L18 7.4q.475.275.738.738t.262.987v5.75q0 .525-.263.988T18 16.6l-5 2.875q-.475.275-1 .275t-1-.275ZM2 7V4q0-.825.588-1.413T4 2h3v2H4v3H2Zm5 15H4q-.825 0-1.413-.588T2 20v-3h2v3h3v2Zm10 0v-2h3v-3h2v3q0 .825-.588 1.413T20 22h-3Zm3-15V4h-3V2h3q.825 0 1.413.588T22 4v3h-2ZM8.05 8.525l-1.05.6v1.125l4 2.325v4.6l1 .575l1-.575v-4.6l4-2.325V9.125l-1.05-.6L12 10.85L8.05 8.525Z"/>`),
		g.Group(children),
	)
}