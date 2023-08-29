package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 12A6 6 0 1 0 6 0a6 6 0 0 0 0 12Zm.5-8.75v2.405l1.488 1.276a.75.75 0 1 1-.976 1.138l-1.75-1.5A.75.75 0 0 1 5 6V3.25a.75.75 0 0 1 1.5 0Z"/>`),
		g.Group(children),
	)
}