package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-6q0-.675.275-1.225T3 10.8V8q0-1.25.875-2.125T6 5h4q.575 0 1.075.213T12 5.8q.425-.375.925-.587T14 5h4q1.25 0 2.125.875T21 8v2.8q.45.425.725.975T22 13v6h-2v-2H4v2H2Zm11-9h6V8q0-.425-.288-.713T18 7h-4q-.425 0-.713.288T13 8v2Zm-8 0h6V8q0-.425-.288-.713T10 7H6q-.425 0-.713.288T5 8v2Z"/>`),
		g.Group(children),
	)
}