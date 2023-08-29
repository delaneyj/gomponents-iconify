package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolunteerActivism(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13q-1.15-1.05-2.238-2.063T11.837 8.95Q11 7.975 10.5 7.063T10 5.3q0-1.4.95-2.35T13.3 2q.8 0 1.5.338t1.2.912q.5-.575 1.2-.912T18.7 2q1.4 0 2.35.95T22 5.3q0 .85-.5 1.763T20.163 8.95q-.838.975-1.913 1.988T16 13Zm-2 9.5l-7-1.95V11h1.95l6.2 2.3q.825.3 1.337 1.05T17 16h-2q-1.05 0-1.65-.075T12.3 15.7l-2-.65l-.3.95l1.575.575q.7.275 1.275.35T14.2 17H19q1.65 0 2.325.537T22 19v1l-8 2.5ZM1 22V11h4v11H1Z"/>`),
		g.Group(children),
	)
}