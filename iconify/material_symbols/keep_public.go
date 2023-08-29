package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeepPublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 23l-1-1v-6H3v-2l2-2V5H4V3h7.125Q10.1 4 9.55 5.288T9 8q0 2.625 1.713 4.588T15 14.925V16h-5v6l-1 1Zm7-10q-2.075 0-3.538-1.463T11 8q0-2.075 1.463-3.538T16 3q2.075 0 3.538 1.463T21 8q0 2.075-1.463 3.538T16 13Zm0-5q.625 0 1.063-.438T17.5 6.5q0-.625-.438-1.063T16 5q-.625 0-1.063.438T14.5 6.5q0 .625.438 1.063T16 8Zm0 3q.75 0 1.4-.35t1.075-.975q-.575-.35-1.2-.513T16 9q-.65 0-1.275.163t-1.2.512q.425.625 1.075.975T16 11Z"/>`),
		g.Group(children),
	)
}