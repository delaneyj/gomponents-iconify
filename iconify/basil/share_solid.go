package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.5a3 3 0 0 0-2.902 3.765a.75.75 0 0 0-.207.077l-2.757 1.503L8.128 9.85a.755.755 0 0 0-.1.068a3 3 0 1 0 .682 4.612l2.926 1.627l2.954 1.611a3 3 0 1 0 .719-1.316l-2.947-1.608l-2.946-1.636a3.007 3.007 0 0 0-.308-2.19l3.258-1.862l2.743-1.497a.75.75 0 0 0 .177-.133A3 3 0 1 0 17.5 2.5Z"/>`),
		g.Group(children),
	)
}