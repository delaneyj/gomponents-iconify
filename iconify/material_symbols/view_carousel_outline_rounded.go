package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewCarouselOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17q-.425 0-.713-.288T2 16V8q0-.425.288-.713T3 7h2q.425 0 .713.288T6 8v8q0 .425-.288.713T5 17H3Zm5 2q-.425 0-.713-.288T7 18V6q0-.425.288-.713T8 5h8q.425 0 .713.288T17 6v12q0 .425-.288.713T16 19H8ZM9 7v10V7Zm10 10q-.425 0-.713-.288T18 16V8q0-.425.288-.713T19 7h2q.425 0 .713.288T22 8v8q0 .425-.288.713T21 17h-2ZM9 7v10h6V7H9Z"/>`),
		g.Group(children),
	)
}