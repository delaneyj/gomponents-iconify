package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelvesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 23q-.425 0-.713-.288T3 22V2q0-.425.288-.713T4 1q.425 0 .713.288T5 2v1h14V2q0-.425.288-.713T20 1q.425 0 .713.288T21 2v20q0 .425-.288.713T20 23q-.425 0-.713-.288T19 22v-1H5v1q0 .425-.288.713T4 23Zm1-12h2V8q0-.425.288-.713T8 7h4q.425 0 .713.288T13 8v3h6V5H5v6Zm0 8h6v-3q0-.425.288-.713T12 15h4q.425 0 .713.288T17 16v3h2v-6H5v6Zm4-8h2V9H9v2Zm4 8h2v-2h-2v2Zm-4-8h2h-2Zm4 8h2h-2Z"/>`),
		g.Group(children),
	)
}