package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12ZM4 9q-.425 0-.713-.288T3 8V6q0-.425.288-.713T4 5h2q.425 0 .713.288T7 6v2q0 .425-.288.713T6 9H4Zm5 0q-.425 0-.713-.288T8 8V6q0-.425.288-.713T9 5h11q.425 0 .713.288T21 6v2q0 .425-.288.713T20 9H9Zm0 5q-.425 0-.713-.288T8 13v-2q0-.425.288-.713T9 10h11q.425 0 .713.288T21 11v2q0 .425-.288.713T20 14H9Zm0 5q-.425 0-.713-.288T8 18v-2q0-.425.288-.713T9 15h11q.425 0 .713.288T21 16v2q0 .425-.288.713T20 19H9Zm-5 0q-.425 0-.713-.288T3 18v-2q0-.425.288-.713T4 15h2q.425 0 .713.288T7 16v2q0 .425-.288.713T6 19H4Zm0-5q-.425 0-.713-.288T3 13v-2q0-.425.288-.713T4 10h2q.425 0 .713.288T7 11v2q0 .425-.288.713T6 14H4Z"/>`),
		g.Group(children),
	)
}