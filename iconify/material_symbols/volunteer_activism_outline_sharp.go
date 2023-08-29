package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolunteerActivismOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13q-2.35-2.125-4.175-4.063T10 5.3q0-1.4.95-2.35T13.3 2q.8 0 1.5.338t1.2.912q.5-.575 1.2-.912T18.7 2q1.4 0 2.35.95T22 5.3q0 1.7-1.825 3.638T16 13Zm0-2.7q1.475-1.4 2.738-2.788T20 5.3q0-.575-.363-.937T18.7 4q-.35 0-.663.138t-.537.412L16 6.35l-1.5-1.8q-.225-.275-.537-.413T13.3 4q-.575 0-.938.363T12 5.3q0 .825 1.263 2.213T16 10.3Zm0-3.15ZM14 22.5l-7-1.95V22H1V11h7.95L17 14v2h5v4l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85V18h-7.075L9.7 16.95l.6-1.9l2.925.95H15v-.65L8.6 13H7v5.5l6.95 1.9Z"/>`),
		g.Group(children),
	)
}