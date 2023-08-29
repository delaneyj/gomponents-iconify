package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1a2 2 0 0 0-2 2v10a2 2 0 1 0 4 0V3a2 2 0 0 0-2-2Zm1 2v10a1 1 0 1 1-2 0V3a1 1 0 0 1 2 0ZM3 7a2 2 0 0 0-2 2v4a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Zm1 6a1 1 0 1 1-2 0V9a1 1 0 0 1 2 0v4Zm7-6a2 2 0 1 1 4 0v6a2 2 0 1 1-4 0V7Zm3 0a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0V7Z"/>`),
		g.Group(children),
	)
}