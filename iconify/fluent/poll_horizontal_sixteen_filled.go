package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollHorizontalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 6a2 2 0 1 1 0 4H3a2 2 0 1 1 0-4h10ZM7 1a2 2 0 1 1 0 4H3a2 2 0 1 1 0-4h4Zm2 10a2 2 0 1 1 0 4H3a2 2 0 1 1 0-4h6Z"/>`),
		g.Group(children),
	)
}