package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.7 2.86a1 1 0 0 0-.84-.2a3 3 0 0 1-2.33-.48a1 1 0 0 0-1.15 0a3 3 0 0 1-2.33.48a1 1 0 0 0-.84.2a1 1 0 0 0-.37.77V7a4.56 4.56 0 0 0 1.91 3.7l1.62 1.16a1 1 0 0 0 1.17 0l1.62-1.16A4.56 4.56 0 0 0 22.07 7V3.63a1 1 0 0 0-.37-.77ZM20.07 7A2.57 2.57 0 0 1 19 9l-1 .74L16.91 9a2.57 2.57 0 0 1-1.07-2V4.72A5.17 5.17 0 0 0 18 4.17a5.12 5.12 0 0 0 2.11.55Zm-1.14 7a1 1 0 0 0-1.21.72A7 7 0 0 1 10.93 20H5.35l.65-.63A1 1 0 0 0 6 18a7 7 0 0 1 4.93-12a1 1 0 0 0 0-2a9 9 0 0 0-7 14.62l-1.7 1.67a1 1 0 0 0 .7 1.71h8a9 9 0 0 0 8.72-6.75a1 1 0 0 0-.72-1.25Z"/>`),
		g.Group(children),
	)
}