package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.6 12L3.5 6.9L4.2 5l15.6 5.7q.5.2.85.7t.35 1.1V18q0 .825-.587 1.413T19 20H5q-.825 0-1.413-.588T3 18v-4q0-.825.588-1.413T5 12h12.6Zm1.4 6v-4H5v4h14Zm-9-1h8v-2h-8v2Zm-3 0q.425 0 .713-.288T8 16q0-.425-.288-.713T7 15q-.425 0-.713.288T6 16q0 .425.288.713T7 17Zm-2 1v-4v4Z"/>`),
		g.Group(children),
	)
}