package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintConnectOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10h16H4Zm2 10v-3H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.45-.25-.95-.425t-1.05-.25V11q0-.425-.288-.713T19 10H5q-.425 0-.713.288T4 11v4h2v-1q0-.425.288-.713T7 13h7.5q-.4.425-.725.925T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H7q-.425 0-.713-.288T6 20Zm11.95-2.65l2.825-2.825q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712l-3.525 3.525q-.3.3-.7.3t-.7-.3l-1.425-1.425q-.3-.3-.3-.7t.3-.7q.3-.275.7-.275t.7.275l.725.7ZM6 8V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v3h-2V5H8v3H6Z"/>`),
		g.Group(children),
	)
}