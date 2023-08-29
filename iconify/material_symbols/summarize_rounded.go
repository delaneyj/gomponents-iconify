package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SummarizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9q.425 0 .713-.288T9 8q0-.425-.288-.713T8 7q-.425 0-.713.288T7 8q0 .425.288.713T8 9Zm0 4q.425 0 .713-.288T9 12q0-.425-.288-.713T8 11q-.425 0-.713.288T7 12q0 .425.288.713T8 13Zm0 4q.425 0 .713-.288T9 16q0-.425-.288-.713T8 15q-.425 0-.713.288T7 16q0 .425.288.713T8 17Zm-3 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h10.175q.4 0 .763.15t.637.425l3.85 3.85q.275.275.425.638t.15.762V19q0 .825-.588 1.413T19 21H5ZM15 5v3q0 .425.288.713T16 9h3l-4-4Z"/>`),
		g.Group(children),
	)
}