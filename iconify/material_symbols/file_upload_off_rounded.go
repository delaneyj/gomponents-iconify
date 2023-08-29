package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUploadOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7.85v2.3l-2-2l-1.575-1.575L11.3 4.7q.15-.15.325-.213T12 4.425q.2 0 .375.063t.325.212l3.6 3.6q.3.3.288.7t-.288.7q-.3.3-.713.313t-.712-.288L13 7.85ZM11 11l2 2v2q0 .425-.288.713T12 16q-.425 0-.713-.288T11 15v-4Zm6.15 9H6q-.825 0-1.413-.588T4 18v-2q0-.425.288-.713T5 15q.425 0 .713.288T6 16v2h9.15L2.075 4.925q-.3-.3-.3-.712t.3-.713q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L17.15 20ZM20 17.15l-1.775-1.775q.125-.15.313-.263T19 15q.425 0 .713.288T20 16v1.15Z"/>`),
		g.Group(children),
	)
}