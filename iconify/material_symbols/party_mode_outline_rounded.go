package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartyModeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm0-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Zm0 4.5q1.875 0 3.188-1.313T16.5 13q0-.125-.013-.25t-.037-.25h-2q.05.125.05.25V13q0 1.05-.725 1.775T12 15.5H8.25q.6.9 1.588 1.45T12 17.5Zm-4.45-4h2q-.05-.125-.05-.25V13q0-1.05.725-1.775T12 10.5h3.75q-.6-.9-1.588-1.45T12 8.5q-1.875 0-3.188 1.313T7.5 13q0 .125.013.25t.037.25Z"/>`),
		g.Group(children),
	)
}