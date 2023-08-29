package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 3a2 2 0 1 1 4 0v10a2 2 0 1 1-4 0V3ZM1 9a2 2 0 1 1 4 0v4a2 2 0 1 1-4 0V9Zm10-2a2 2 0 1 1 4 0v6a2 2 0 1 1-4 0V7Z"/>`),
		g.Group(children),
	)
}