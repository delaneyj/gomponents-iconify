package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MissedVideoCallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.7 13.15L8.5 11H9q.425 0 .713-.288T10 10q0-.425-.288-.713T9 9H6q-.425 0-.713.288T5 10v3q0 .425.288.713T6 14q.425 0 .713-.288T7 13v-.7l3 3q.125.125.313.2t.387.075q.2 0 .388-.075t.312-.225l3.125-3.15q.275-.275.275-.675t-.3-.7q-.275-.275-.7-.275t-.7.275l-2.4 2.4ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.225-.225.537-.113T22 7.7v8.6q0 .35-.313.463t-.537-.113L18 13.5V18q0 .825-.588 1.413T16 20H4Z"/>`),
		g.Group(children),
	)
}