package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandContentRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h3q.425 0 .713.288T11 18q0 .425-.288.713T10 19H6q-.425 0-.713-.288T5 18v-4q0-.425.288-.713T6 13q.425 0 .713.288T7 14v3ZM17 7h-3q-.425 0-.713-.288T13 6q0-.425.288-.713T14 5h4q.425 0 .713.288T19 6v4q0 .425-.288.713T18 11q-.425 0-.713-.288T17 10V7Z"/>`),
		g.Group(children),
	)
}