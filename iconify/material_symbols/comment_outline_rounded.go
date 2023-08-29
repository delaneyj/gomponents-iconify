package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14h10q.425 0 .713-.288T18 13q0-.425-.288-.713T17 12H7q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-3h10q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9H7q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0-3h10q.425 0 .713-.288T18 7q0-.425-.288-.713T17 6H7q-.425 0-.713.288T6 7q0 .425.288.713T7 8ZM4 18q-.825 0-1.413-.588T2 16V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v15.575q0 .675-.613.938T20.3 20.3L18 18H4Zm0-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}