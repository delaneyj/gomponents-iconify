package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewArrayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19q-.425 0-.713-.288T18 18V6q0-.425.288-.713T19 5h1q.425 0 .713.288T21 6v12q0 .425-.288.713T20 19h-1ZM8 19q-.425 0-.713-.288T7 18V6q0-.425.288-.713T8 5h8q.425 0 .713.288T17 6v12q0 .425-.288.713T16 19H8Zm-4 0q-.425 0-.713-.288T3 18V6q0-.425.288-.713T4 5h1q.425 0 .713.288T6 6v12q0 .425-.288.713T5 19H4Z"/>`),
		g.Group(children),
	)
}