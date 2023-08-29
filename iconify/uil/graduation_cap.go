package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.49 10.19l-1-.55l-9-5h-.11a1.06 1.06 0 0 0-.19-.06h-.37a1.17 1.17 0 0 0-.2.06h-.11l-9 5a1 1 0 0 0 0 1.74L4 12.76v4.74a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-4.74l2-1.12v2.86a1 1 0 0 0 2 0v-3.44a1 1 0 0 0-.51-.87ZM16 17.5a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3.63l4.51 2.5l.15.06h.09a1 1 0 0 0 .25 0a1 1 0 0 0 .25 0h.09a.47.47 0 0 0 .15-.06l4.51-2.5Zm-5-3.14L4.06 10.5L11 6.64l6.94 3.86Z"/>`),
		g.Group(children),
	)
}