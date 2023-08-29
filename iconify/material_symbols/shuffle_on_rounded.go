package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23q-.825 0-1.413-.588T1 21V3q0-.825.588-1.413T3 1h18q.825 0 1.413.588T23 3v18q0 .825-.588 1.413T21 23H3Zm6.175-12.425l1.4-1.4L6.1 4.7q-.275-.275-.7-.275t-.7.275q-.275.275-.275.7t.275.7l4.475 4.475ZM15 20h4q.425 0 .713-.288T20 19v-4q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1.55l-3.175-3.175L13.4 14.8l3.2 3.2H15q-.425 0-.712.288T14 19q0 .425.288.713T15 20Zm-10.3-.7q.275.275.7.275t.7-.275L18 7.4V9q0 .425.288.713T19 10q.425 0 .713-.288T20 9V5q0-.425-.288-.713T19 4h-4q-.425 0-.713.288T14 5q0 .425.288.713T15 6h1.6L4.7 17.9q-.275.275-.275.7t.275.7Z"/>`),
		g.Group(children),
	)
}