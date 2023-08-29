package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSnippetOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14V9.825L14.175 5H5v14Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h9.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V19q0 .825-.588 1.413T19 21H5Zm3-4h8q.425 0 .713-.288T17 16q0-.425-.288-.713T16 15H8q-.425 0-.713.288T7 16q0 .425.288.713T8 17Zm0-4h8q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13Zm0-4h5q.425 0 .713-.288T14 8q0-.425-.288-.713T13 7H8q-.425 0-.713.288T7 8q0 .425.288.713T8 9ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}