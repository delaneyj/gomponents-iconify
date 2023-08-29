package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewInArOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19.475L6 16.6q-.475-.275-.738-.725t-.262-1v-5.75q0-.55.263-1T6 7.4l5-2.875q.475-.275 1-.275t1 .275L18 7.4q.475.275.738.725t.262 1v5.75q0 .55-.263 1T18 16.6l-5 2.875q-.475.275-1 .275t-1-.275Zm0-2.3v-4.6L7 10.25v4.625l4 2.3Zm2 0l4-2.3V10.25l-4 2.325v4.6ZM2 7V4q0-.825.588-1.413T4 2h3v2H4v3H2Zm5 15H4q-.825 0-1.413-.588T2 20v-3h2v3h3v2Zm10 0v-2h3v-3h2v3q0 .825-.588 1.413T20 22h-3Zm3-15V4h-3V2h3q.825 0 1.413.588T22 4v3h-2Zm-8 3.85l3.95-2.325L12 6.25L8.05 8.525L12 10.85Zm0 1.125Zm0-1.125Zm1 1.725Zm-2 0Z"/>`),
		g.Group(children),
	)
}