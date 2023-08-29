package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignFlexCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21v-7.5H4q-.425 0-.713-.288T3 12.5v-1q0-.425.288-.713T4 10.5h7V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v7.5h7q.425 0 .713.288T21 11.5v1q0 .425-.288.713T20 13.5h-7V21q0 .425-.288.713T12 22Z"/>`),
		g.Group(children),
	)
}