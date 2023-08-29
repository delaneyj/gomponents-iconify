package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7H9q-.425 0-.713-.288T8 6q0-.425.288-.713T9 5h9q.425 0 .713.288T19 6v9q0 .425-.288.713T18 16q-.425 0-.713-.288T17 15V7Zm-5 5H4q-.425 0-.713-.288T3 11q0-.425.288-.713T4 10h9q.425 0 .713.288T14 11v9q0 .425-.288.713T13 21q-.425 0-.713-.288T12 20v-8Z"/>`),
		g.Group(children),
	)
}