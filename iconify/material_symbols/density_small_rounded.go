package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DensitySmallRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.425 0-.713-.288T3 21q0-.425.288-.713T4 20h16q.425 0 .713.288T21 21q0 .425-.288.713T20 22H4Zm0-6q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h16q.425 0 .713.288T21 15q0 .425-.288.713T20 16H4Zm0-6q-.425 0-.713-.288T3 9q0-.425.288-.713T4 8h16q.425 0 .713.288T21 9q0 .425-.288.713T20 10H4Zm0-6q-.425 0-.713-.288T3 3q0-.425.288-.713T4 2h16q.425 0 .713.288T21 3q0 .425-.288.713T20 4H4Z"/>`),
		g.Group(children),
	)
}