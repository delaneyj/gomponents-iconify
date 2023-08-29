package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 20v-3.5q0-.425.288-.713T7.5 15.5H11V12q0-.425.288-.713T12 11h3.5V7.5q0-.425.288-.713T16.5 6.5H20V4q0-.425.288-.713T21 3q.425 0 .713.288T22 4v3.5q0 .425-.288.713T21 8.5h-3.5V12q0 .425-.288.713T16.5 13H13v3.5q0 .425-.288.713T12 17.5H8.5V21q0 .425-.288.713T7.5 22H4q-.425 0-.713-.288T3 21q0-.425.288-.713T4 20h2.5Z"/>`),
		g.Group(children),
	)
}