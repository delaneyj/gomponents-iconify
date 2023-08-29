package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentBankRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm6.5-14v6.125q0 .3.25.438t.5-.013l1.375-.825q.2-.125.375-.125t.375.125l1.375.825q.25.15.5.013t.25-.438V4h-5Z"/>`),
		g.Group(children),
	)
}