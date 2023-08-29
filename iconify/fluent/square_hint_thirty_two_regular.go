package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareHintThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 3a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4Zm0 24a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4ZM4 19a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1Zm23-1a1 1 0 1 0 2 0v-4a1 1 0 1 0-2 0v4ZM9 4a1 1 0 0 0-1-1h-.25A4.75 4.75 0 0 0 3 7.75V8a1 1 0 0 0 2 0v-.25A2.75 2.75 0 0 1 7.75 5H8a1 1 0 0 0 1-1ZM8 29a1 1 0 1 0 0-2h-.25A2.75 2.75 0 0 1 5 24.25V24a1 1 0 1 0-2 0v.25A4.75 4.75 0 0 0 7.75 29H8ZM23 4a1 1 0 0 1 1-1h.25A4.75 4.75 0 0 1 29 7.75V8a1 1 0 1 1-2 0v-.25A2.75 2.75 0 0 0 24.25 5H24a1 1 0 0 1-1-1Zm1 25a1 1 0 1 1 0-2h.25A2.75 2.75 0 0 0 27 24.25V24a1 1 0 1 1 2 0v.25A4.75 4.75 0 0 1 24.25 29H24Z"/>`),
		g.Group(children),
	)
}