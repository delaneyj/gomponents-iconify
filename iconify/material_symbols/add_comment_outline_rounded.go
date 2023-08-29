package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddCommentOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11v2q0 .425.288.713T12 14q.425 0 .713-.288T13 13v-2h2q.425 0 .713-.288T16 10q0-.425-.288-.713T15 9h-2V7q0-.425-.288-.713T12 6q-.425 0-.713.288T11 7v2H9q-.425 0-.713.288T8 10q0 .425.288.713T9 11h2Zm-5 7l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm-2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}