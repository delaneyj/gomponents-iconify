package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartyMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L9 3h6l1.85 2H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm8-3.5q1.875 0 3.188-1.313T16.5 13q0-.125-.013-.25t-.037-.25h-2q.05.125.05.25V13q0 1.05-.725 1.775T12 15.5H8.25q.6.9 1.588 1.45T12 17.5Zm-4.45-4h2q-.05-.125-.05-.25V13q0-1.05.725-1.775T12 10.5h3.75q-.6-.9-1.588-1.45T12 8.5q-1.875 0-3.188 1.313T7.5 13q0 .125.013.25t.037.25Z"/>`),
		g.Group(children),
	)
}