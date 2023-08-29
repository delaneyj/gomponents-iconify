package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionStraightRightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H4a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm-8 21l-1.414-1.414L23.172 17H6v-2h17.172l-4.586-4.586L20 9l7 7Z"/><path fill="none" d="m20 23l-1.414-1.414L23.172 17H6v-2h17.172l-4.586-4.586L20 9l7 7Z"/>`),
		g.Group(children),
	)
}