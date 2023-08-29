package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMoreUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19q-.425 0-.713-.288T12 18v-8H4q-.425 0-.713-.288T3 9q0-.425.288-.713T4 8h9q.425 0 .713.288T14 9v9q0 .425-.288.713T13 19Zm5-5q-.425 0-.713-.288T17 13V5H9q-.425 0-.713-.288T8 4q0-.425.288-.713T9 3h9q.425 0 .713.288T19 4v9q0 .425-.288.713T18 14Z"/>`),
		g.Group(children),
	)
}