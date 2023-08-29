package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalletRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v1.25q0 .425-.288.713T4 22.25H3q-.425 0-.713-.288T2 21.25V19q0-.425.288-.713T3 18h18q.425 0 .713.288T22 19v2.25q0 .425-.288.713T21 22.25h-1q-.425 0-.713-.288T19 21.25V20h-5.5v1.25q0 .425-.288.713t-.712.287h-1q-.425 0-.713-.288t-.287-.712V20H5Zm1-4q-.425 0-.713-.288T5 15V3q0-.425.288-.713T6 2h12q.425 0 .713.288T19 3v12q0 .425-.288.713T18 16H6Zm8-8q.425 0 .713-.288T15 7q0-.425-.288-.713T14 6h-4q-.425 0-.713.288T9 7q0 .425.288.713T10 8h4Z"/>`),
		g.Group(children),
	)
}