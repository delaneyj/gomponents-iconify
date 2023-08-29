package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NorthEastRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.7 19.3q-.275-.275-.275-.7t.275-.7L15.6 7H10q-.425 0-.713-.288T9 6q0-.425.288-.713T10 5h8q.425 0 .713.288T19 6v8q0 .425-.288.713T18 15q-.425 0-.713-.288T17 14V8.4L6.1 19.3q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}