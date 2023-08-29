package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolunteerActivismSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13q-2.35-2.125-4.175-4.063T10 5.3q0-1.4.95-2.35T13.3 2q.8 0 1.5.338t1.2.912q.5-.575 1.2-.912T18.7 2q1.4 0 2.35.95T22 5.3q0 1.7-1.825 3.638T16 13ZM1 22V11h4v11H1Zm13 0l-7-1.975V11h1.975L17 14v2h-4l-1.75-.675l-.35.925L13 17h9v2l-8 3Z"/>`),
		g.Group(children),
	)
}