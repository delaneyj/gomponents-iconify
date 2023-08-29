package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanZoomRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.4 19H8q.425 0 .713.288T9 20q0 .425-.288.713T8 21H4q-.425 0-.713-.288T3 20v-4q0-.425.288-.713T4 15q.425 0 .713.288T5 16v1.6l2.4-2.4q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L6.4 19ZM19 6.4l-2.4 2.4q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7L17.6 5H16q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h4q.425 0 .713.288T21 4v4q0 .425-.288.713T20 9q-.425 0-.713-.288T19 8V6.4Z"/>`),
		g.Group(children),
	)
}