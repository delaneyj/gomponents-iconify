package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintConnectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.95 17.35l2.825-2.825q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712l-3.525 3.525q-.3.3-.7.3t-.7-.3l-1.425-1.425q-.3-.3-.3-.7t.3-.7q.3-.275.7-.275t.7.275l.725.7ZM7 21q-.425 0-.713-.288T6 20v-3H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.675-.35-1.413-.55t-1.512-.2q-1.95 0-3.538 1.1T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H7ZM6 7V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v2H6Z"/>`),
		g.Group(children),
	)
}