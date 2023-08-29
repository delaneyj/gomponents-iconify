package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarBadgeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 16c-1.9 0-3.5 1.6-3.5 3.5s1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5s-1.6-3.5-3.5-3.5M14 19.5c0-.17 0-.33.03-.5H5V9h14v5.03c.17-.03.33-.03.5-.03c.5 0 1 .08 1.5.21V5a2 2 0 0 0-2-2h-1V1h-2v2H8V1H6v2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h9.21c-.13-.5-.21-1-.21-1.5M5 5h14v2H5V5Z"/>`),
		g.Group(children),
	)
}