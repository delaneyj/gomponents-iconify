package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecklistRtlRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.35 16.175l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.275.7t-.275.7l-4.225 4.25q-.3.3-.7.3t-.7-.3l-2.15-2.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.425 1.425Zm0-8l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.275.7t-.275.7l-4.225 4.25q-.3.3-.7.3t-.7-.3l-2.15-2.15q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.425 1.425ZM3 17q-.425 0-.713-.288T2 16q0-.425.288-.713T3 15h7q.425 0 .713.288T11 16q0 .425-.288.713T10 17H3Zm0-8q-.425 0-.713-.288T2 8q0-.425.288-.713T3 7h7q.425 0 .713.288T11 8q0 .425-.288.713T10 9H3Z"/>`),
		g.Group(children),
	)
}