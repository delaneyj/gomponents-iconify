package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8Zm7-2.622c0-.755-.99-1.037-1.387-.396L5.07 9.084a.6.6 0 0 0 .51.916H8v.5a.5.5 0 0 0 1 0V10h.5a.5.5 0 0 0 0-1H9V5.378Zm-1 .877V9H6.3L8 6.255Z"/>`),
		g.Group(children),
	)
}