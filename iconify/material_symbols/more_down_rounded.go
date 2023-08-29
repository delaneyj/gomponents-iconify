package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h8q.425 0 .713.288T16 18q0 .425-.288.713T15 19H6q-.425 0-.713-.288T5 18V9q0-.425.288-.713T6 8q.425 0 .713.288T7 9v8Zm5-5h8q.425 0 .713.288T21 13q0 .425-.288.713T20 14h-9q-.425 0-.713-.288T10 13V4q0-.425.288-.713T11 3q.425 0 .713.288T12 4v8Z"/>`),
		g.Group(children),
	)
}