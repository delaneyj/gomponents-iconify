package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.808 1.394l19.799 19.798l-1.415 1.415L17.585 19H6.455L2 22.5V4c0-.17.042-.329.116-.468l-.722-.724l1.414-1.414ZM21 3a1 1 0 0 1 1 1v13.785L7.214 3H21Z"/>`),
		g.Group(children),
	)
}