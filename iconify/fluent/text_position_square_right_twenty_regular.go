package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextPositionSquareRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.5 3a.5.5 0 0 1 0 1h-13a.5.5 0 0 1 0-1h13ZM8 6a.5.5 0 0 1 0 1H3.5a.5.5 0 0 1 0-1H8Zm.5 3.5A.5.5 0 0 0 8 9H3.5a.5.5 0 0 0 0 1H8a.5.5 0 0 0 .5-.5ZM8 12a.5.5 0 0 1 0 1H3.5a.5.5 0 0 1 0-1H8Zm9 3.5a.5.5 0 0 0-.5-.5h-13a.5.5 0 0 0 0 1h13a.5.5 0 0 0 .5-.5Zm-1-7a2.5 2.5 0 0 0-5 0v5a.5.5 0 0 1-1 0v-5a3.5 3.5 0 1 1 7 0v5a.5.5 0 0 1-1 0v-5Z"/>`),
		g.Group(children),
	)
}