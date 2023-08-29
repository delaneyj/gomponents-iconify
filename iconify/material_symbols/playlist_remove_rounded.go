package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistRemoveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 19.4l-1.9 1.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l1.9-1.9l-1.9-1.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L18.4 18l1.9 1.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L17 19.4ZM4 16q-.425 0-.713-.288T3 15q0-.425.288-.713T4 14h5q.425 0 .713.288T10 15q0 .425-.288.713T9 16H4Zm0-4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h9q.425 0 .713.288T14 11q0 .425-.288.713T13 12H4Zm0-4q-.425 0-.713-.288T3 7q0-.425.288-.713T4 6h9q.425 0 .713.288T14 7q0 .425-.288.713T13 8H4Z"/>`),
		g.Group(children),
	)
}