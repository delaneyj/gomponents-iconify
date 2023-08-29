package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteIt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.038 12.488a1.842 1.842 0 0 0-1.842-1.842H8.804a1.842 1.842 0 0 0-1.842 1.842v23.024a1.842 1.842 0 0 0 1.842 1.842h30.392a1.842 1.842 0 0 0 1.842-1.842Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m38.06 10.646l-.582-2.176a1.842 1.842 0 0 0-2.256-1.302l-12.98 3.478M41.038 33.26l1.096-.294a1.842 1.842 0 0 0 1.303-2.256l-2.4-8.953M9.94 37.354l.582 2.175a1.842 1.842 0 0 0 2.256 1.303l12.98-3.478M6.962 14.74l-1.096.294a1.842 1.842 0 0 0-1.303 2.256l2.414 9.008m5.971-10.126h22.103m-22.103 5.219h22.103m-22.103 5.218h22.103m-22.103 5.219H24"/>`),
		g.Group(children),
	)
}