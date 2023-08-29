package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalletOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-4h20v4h-3v-2h-5.5v2h-3v-2H5v2H2Zm4-6q-.425 0-.713-.288T5 15V3q0-.425.288-.713T6 2h12q.425 0 .713.288T19 3v12q0 .425-.288.713T18 16H6Zm1-2h10V4H7v10Zm2-6h6V6H9v2Zm-2 6V4v10Z"/>`),
		g.Group(children),
	)
}