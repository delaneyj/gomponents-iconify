package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 9V7h14v2H7Zm0 4v-2h14v2H7Zm0 4v-2h14v2H7ZM4 9q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7q.425 0 .713.288T5 8q0 .425-.288.713T4 9Zm0 4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11q.425 0 .713.288T5 12q0 .425-.288.713T4 13Zm0 4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15q.425 0 .713.288T5 16q0 .425-.288.713T4 17Z"/>`),
		g.Group(children),
	)
}