package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionRightTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM8 6h2v8H8Zm11 18l-1.414-1.414L22.172 18H10v8H8v-8a2 2 0 0 1 2-2h12.172l-4.586-4.586L19 10l7 7Z"/><path fill="none" d="M8 6h2v8H8Zm11 18l-1.414-1.414L22.172 18H10v8H8v-8a2 2 0 0 1 2-2h12.172l-4.586-4.586L19 10l7 7Z"/>`),
		g.Group(children),
	)
}