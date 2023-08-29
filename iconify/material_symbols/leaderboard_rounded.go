package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeaderboardRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.425 0-.713-.288T2 20V10q0-.425.288-.713T3 9h3.5q.425 0 .713.288T7.5 10v10q0 .425-.288.713T6.5 21H3Zm7.25 0q-.425 0-.713-.288T9.25 20V4q0-.425.288-.713T10.25 3h3.5q.425 0 .713.288T14.75 4v16q0 .425-.288.713T13.75 21h-3.5Zm7.25 0q-.425 0-.713-.288T16.5 20v-8q0-.425.288-.713T17.5 11H21q.425 0 .713.288T22 12v8q0 .425-.288.713T21 21h-3.5Z"/>`),
		g.Group(children),
	)
}