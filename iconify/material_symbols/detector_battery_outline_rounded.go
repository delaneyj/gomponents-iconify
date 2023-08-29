package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorBatteryOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h8v-3H9v3Zm-3 2q-.425 0-.713-.288T5 20v-5q0-.425.288-.713T6 14h12q.425 0 .713.288T19 15v1h1q.425 0 .713.288T21 17v1q0 .425-.288.713T20 19h-1v1q0 .425-.288.713T18 21H6ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}