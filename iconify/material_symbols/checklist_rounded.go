package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecklistRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.525 16.175l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.275.7t-.275.7L6.25 18.3q-.3.3-.7.3t-.7-.3L2.7 16.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.425 1.425Zm0-8l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.275.7t-.275.7L6.25 10.3q-.3.3-.7.3t-.7-.3L2.7 8.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.425 1.425ZM14 17q-.425 0-.713-.288T13 16q0-.425.288-.713T14 15h7q.425 0 .713.288T22 16q0 .425-.288.713T21 17h-7Zm0-8q-.425 0-.713-.288T13 8q0-.425.288-.713T14 7h7q.425 0 .713.288T22 8q0 .425-.288.713T21 9h-7Z"/>`),
		g.Group(children),
	)
}