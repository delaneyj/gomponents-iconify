package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20q0-.425.288-.713T7 19h10q.425 0 .713.288T18 20q0 .425-.288.713T17 21H7Z"/>`),
		g.Group(children),
	)
}