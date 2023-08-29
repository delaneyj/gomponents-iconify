package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeScannerRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7q-.425 0-.713-.288T2 6V3q0-.425.288-.713T3 2h3q.425 0 .713.288T7 3q0 .425-.288.713T6 4H4v2q0 .425-.288.713T3 7Zm0 15q-.425 0-.713-.288T2 21v-3q0-.425.288-.713T3 17q.425 0 .713.288T4 18v2h2q.425 0 .713.288T7 21q0 .425-.288.713T6 22H3Zm15 0q-.425 0-.713-.288T17 21q0-.425.288-.713T18 20h2v-2q0-.425.288-.713T21 17q.425 0 .713.288T22 18v3q0 .425-.288.713T21 22h-3Zm3-15q-.425 0-.713-.288T20 6V4h-2q-.425 0-.713-.288T17 3q0-.425.288-.713T18 2h3q.425 0 .713.288T22 3v3q0 .425-.288.713T21 7Zm-3.5 10.5H19V19h-1.5v-1.5Zm0-3H19V16h-1.5v-1.5ZM16 16h1.5v1.5H16V16Zm-1.5 1.5H16V19h-1.5v-1.5ZM13 16h1.5v1.5H13V16Zm3-3h1.5v1.5H16V13Zm-1.5 1.5H16V16h-1.5v-1.5ZM13 13h1.5v1.5H13V13Zm6-8v6h-6V5h6Zm-8 8v6H5v-6h6Zm0-8v6H5V5h6ZM9.5 17.5v-3h-3v3h3Zm0-8v-3h-3v3h3Zm8 0v-3h-3v3h3Z"/>`),
		g.Group(children),
	)
}