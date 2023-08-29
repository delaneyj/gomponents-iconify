package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewArrayOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17V7v10Zm10 2q-.425 0-.713-.288T18 18V6q0-.425.288-.713T19 5h1q.425 0 .713.288T21 6v12q0 .425-.288.713T20 19h-1ZM8 19q-.425 0-.713-.288T7 18V6q0-.425.288-.713T8 5h8q.425 0 .713.288T17 6v12q0 .425-.288.713T16 19H8Zm-4 0q-.425 0-.713-.288T3 18V6q0-.425.288-.713T4 5h1q.425 0 .713.288T6 6v12q0 .425-.288.713T5 19H4ZM9 7v10h6V7H9Z"/>`),
		g.Group(children),
	)
}