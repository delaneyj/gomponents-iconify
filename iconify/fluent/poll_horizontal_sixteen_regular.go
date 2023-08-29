package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PollHorizontalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 8a2 2 0 0 0-2-2H3a2 2 0 1 0 0 4h10a2 2 0 0 0 2-2Zm-2 1H3a1 1 0 1 1 0-2h10a1 1 0 1 1 0 2ZM9 3a2 2 0 0 0-2-2H3a2 2 0 1 0 0 4h4a2 2 0 0 0 2-2ZM3 4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2H3Zm6 7a2 2 0 1 1 0 4H3a2 2 0 1 1 0-4h6Zm0 3a1 1 0 1 0 0-2H3a1 1 0 1 0 0 2h6Z"/>`),
		g.Group(children),
	)
}