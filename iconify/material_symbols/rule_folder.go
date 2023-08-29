package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleFolder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h6l2 2h8q.825 0 1.413.588T22 8v10q0 .825-.588 1.413T20 20H4Zm3.825-3.625l4.95-4.95L11.35 10l-3.525 3.55L6.4 12.125L5 13.55l2.825 2.825Zm6.575 0l1.6-1.6l1.6 1.6l1.4-1.4l-1.6-1.6l1.6-1.6l-1.4-1.4l-1.6 1.6l-1.6-1.6l-1.4 1.4l1.6 1.6l-1.6 1.6l1.4 1.4Z"/>`),
		g.Group(children),
	)
}