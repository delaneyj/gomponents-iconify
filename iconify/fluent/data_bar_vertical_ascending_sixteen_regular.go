package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalAscendingSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 3a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V3Zm2 11a1 1 0 0 0 1-1V3a1 1 0 1 0-2 0v10a1 1 0 0 0 1 1ZM1 9a2 2 0 1 1 4 0v4a2 2 0 1 1-4 0V9Zm2 5a1 1 0 0 0 1-1V9a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm5-9a2 2 0 0 0-2 2v6a2 2 0 1 0 4 0V7a2 2 0 0 0-2-2Zm0 1a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}