package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraFrontOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.15 16.65L18 13.5v1.675L6.825 4H16q.825 0 1.413.588T18 6v4.5l3.15-3.15q.25-.25.55-.125t.3.475v8.6q0 .35-.3.475t-.55-.125Zm-1.3 6L1.35 4.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.5 18.5q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275ZM4 4l14 14q0 .825-.588 1.413T16 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4Zm2 12h8v-.55q0-1.1-1.1-1.775T10 13q-1.8 0-2.9.675T6 15.45V16Z"/>`),
		g.Group(children),
	)
}