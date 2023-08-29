package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorFrontSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21h18v-2h-2V3H5v16H3v2Zm11-8q-.425 0-.713-.288T13 12q0-.425.288-.713T14 11q.425 0 .713.288T15 12q0 .425-.288.713T14 13Z"/>`),
		g.Group(children),
	)
}