package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAddRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h5q.425 0 .713.288T10 15q0 .425-.288.713T9 16H4Zm0-4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h9q.425 0 .713.288T14 11q0 .425-.288.713T13 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h9q.425 0 .713.288T14 7q0 .425-.288.713T13 8H4Zm13 12q-.425 0-.713-.288T16 19v-3h-3q-.425 0-.713-.288T12 15q0-.425.288-.713T13 14h3v-3q0-.425.288-.713T17 10q.425 0 .713.288T18 11v3h3q.425 0 .713.288T22 15q0 .425-.288.713T21 16h-3v3q0 .425-.288.713T17 20Z"/>`),
		g.Group(children),
	)
}