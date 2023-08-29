package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsVerticalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm5-11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}