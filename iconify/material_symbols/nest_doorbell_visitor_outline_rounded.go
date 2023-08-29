package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestDoorbellVisitorOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 18q-.825 0-1.413-.588T16 16q0-.825.588-1.413T18 14q.825 0 1.413.588T20 16q0 .825-.588 1.413T18 18Zm-4 4v-.575q0-1.1 1.1-1.763T18 19q1.8 0 2.9.663t1.1 1.762V22h-8Zm-2-12q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Zm0 2Zm-2 4q0 .825.588 1.413T12 18q.825 0 1.413-.588T14 16q0-.825-.588-1.413T12 14q-.825 0-1.413.588T10 16Zm2 1q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17ZM7 7q0-2.125 1.438-3.563T12 2q2.125 0 3.563 1.438T17 7v4q0 .425-.288.713T16 12q-.425 0-.713-.288T15 11V7q0-1.275-.9-2.138T12 4q-1.2 0-2.1.863T9 7v10q0 1.05.675 1.875t1.65 1.05q.35.075.588.35t.237.625q0 .5-.375.8t-.875.175q-1.7-.375-2.8-1.712T7 17V7Z"/>`),
		g.Group(children),
	)
}