package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentAltShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.57 12a1 1 0 0 0 .58-.19l1.62-1.16A4.56 4.56 0 0 0 10.68 7V3.63a1 1 0 0 0-.37-.77a1 1 0 0 0-.84-.2a3 3 0 0 1-2.33-.48a1 1 0 0 0-1.14 0a3 3 0 0 1-2.33.48a1 1 0 0 0-1.2 1V7a4.55 4.55 0 0 0 1.9 3.7L6 11.81a.94.94 0 0 0 .57.19ZM4.46 7V4.72a5.16 5.16 0 0 0 2.11-.55a5.12 5.12 0 0 0 2.11.55V7a2.57 2.57 0 0 1-1.07 2l-1 .74L5.53 9a2.57 2.57 0 0 1-1.07-2Zm14.08-1h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72L18 17.27a.94.94 0 0 0-.68-.27H8.54a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .67.27a1.15 1.15 0 0 0 .41-.08a1 1 0 0 0 .59-.92V9a3 3 0 0 0-3.03-3Z"/>`),
		g.Group(children),
	)
}