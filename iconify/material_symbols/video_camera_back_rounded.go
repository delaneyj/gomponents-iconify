package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraBackRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.225-.225.537-.113T22 7.7v8.6q0 .35-.313.463t-.537-.113L18 13.5V18q0 .825-.588 1.413T16 20H4Zm2.025-4h7.95q.325 0 .463-.275t-.063-.525l-2.425-3.175q-.15-.2-.4-.2t-.4.2L9.25 14.5L8.1 13q-.15-.2-.4-.2t-.4.2l-1.675 2.2q-.2.25-.063.525t.463.275Z"/>`),
		g.Group(children),
	)
}