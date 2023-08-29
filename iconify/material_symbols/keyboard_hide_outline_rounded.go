package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardHideOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.65 22.65l-2.8-2.8q-.25-.25-.125-.55T9.2 19h5.6q.35 0 .475.3t-.125.55l-2.8 2.8q-.15.15-.35.15t-.35-.15ZM4 17q-.825 0-1.413-.588T2 15V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v10q0 .825-.588 1.413T20 17H4Zm0-2h16V5H4v10Zm5-1h6q.425 0 .713-.288T16 13q0-.425-.288-.713T15 12H9q-.425 0-.713.288T8 13q0 .425.288.713T9 14Zm-5 1V5v10Zm1-4h2V9H5v2Zm3 0h2V9H8v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2Zm3 0h2V9h-2v2ZM5 8h2V6H5v2Zm3 0h2V6H8v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2Zm3 0h2V6h-2v2Z"/>`),
		g.Group(children),
	)
}