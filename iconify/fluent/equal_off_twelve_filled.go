package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualOffTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.854 1.146a.5.5 0 1 0-.708.708L2.793 3.5H2.25a.75.75 0 0 0 0 1.5h2.043l2 2H2.25a.75.75 0 0 0 0 1.5h5.543l2.353 2.354a.5.5 0 0 0 .708-.708l-9-9ZM9.12 7l1.216 1.216A.75.75 0 0 0 9.75 7h-.628Zm-3.5-3.5L7.12 5h2.63a.75.75 0 0 0 0-1.5H5.621Z"/>`),
		g.Group(children),
	)
}