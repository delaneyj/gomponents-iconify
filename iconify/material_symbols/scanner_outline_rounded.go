package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScannerOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.6 12L4.425 7.225q-.375-.15-.537-.512T3.85 5.95q.125-.4.513-.563t.787-.037L19.8 10.7q.5.2.85.7t.35 1.1V18q0 .825-.587 1.413T19 20H5q-.825 0-1.413-.588T3 18v-4q0-.825.588-1.413T5 12h12.6Zm1.4 6v-4H5v4h14Zm-8-1h6q.425 0 .713-.288T18 16q0-.425-.288-.713T17 15h-6q-.425 0-.713.288T10 16q0 .425.288.713T11 17Zm-4 0q.425 0 .713-.288T8 16q0-.425-.288-.713T7 15q-.425 0-.713.288T6 16q0 .425.288.713T7 17Zm-2 1v-4v4Z"/>`),
		g.Group(children),
	)
}