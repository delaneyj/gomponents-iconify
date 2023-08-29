package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolunteerActivismOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13q-1.15-1.05-2.238-2.063T11.837 8.95Q11 7.975 10.5 7.063T10 5.3q0-1.4.95-2.35T13.3 2q.8 0 1.5.338t1.2.912q.5-.575 1.2-.912T18.7 2q1.4 0 2.35.95T22 5.3q0 .85-.5 1.763T20.163 8.95q-.838.975-1.913 1.988T16 13Zm0-2.7q1.475-1.4 2.738-2.788T20 5.3q0-.575-.363-.937T18.7 4q-.35 0-.663.138t-.537.412L16 6.35l-1.5-1.8q-.225-.275-.537-.413T13.3 4q-.575 0-.938.363T12 5.3q0 .825 1.263 2.213T16 10.3Zm-2 12.2l-7-1.95V22H1V11h7.95l6.2 2.3q.825.3 1.337 1.05T17 16h2q1.25 0 2.125.825T22 19v1l-8 2.5ZM3 20h2v-7H3v7Zm10.95.4l5.95-1.85q-.075-.275-.337-.413T19 18h-4.8q-.775 0-1.4-.1t-1.35-.35l-1.725-.6l.575-1.9l2 .65q.45.15 1.05.225T15 16q0-.275-.162-.525t-.388-.325L8.6 13H7v5.5l6.95 1.9ZM5 16.5Zm10-.5Zm-10 .5Zm2 0Zm9-9.35Z"/>`),
		g.Group(children),
	)
}