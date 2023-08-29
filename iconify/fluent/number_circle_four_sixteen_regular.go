package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 8a5 5 0 1 1 10 0A5 5 0 0 1 3 8Zm5-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2Zm1 3.378c0-.755-.99-1.037-1.387-.396L5.07 9.084a.6.6 0 0 0 .51.916H8v.5a.5.5 0 0 0 1 0V10h.5a.5.5 0 0 0 0-1H9V5.378Zm-1 .877V9H6.3L8 6.255Z"/>`),
		g.Group(children),
	)
}