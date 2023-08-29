package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowInsertRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 8.4V16q0 .425-.288.713T7 17q-.425 0-.713-.288T6 16V6q0-.425.288-.713T7 5h10q.425 0 .713.288T18 6q0 .425-.288.713T17 7H9.4l8.9 8.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L8 8.4Z"/>`),
		g.Group(children),
	)
}