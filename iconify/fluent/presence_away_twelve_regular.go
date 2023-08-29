package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 5.655V3.25a.75.75 0 0 0-1.5 0V6a.75.75 0 0 0 .262.57l1.75 1.5a.75.75 0 1 0 .976-1.14L6.5 5.656ZM12 6A6 6 0 1 1 0 6a6 6 0 0 1 12 0Zm-1.5 0a4.5 4.5 0 1 0-9 0a4.5 4.5 0 0 0 9 0Z"/>`),
		g.Group(children),
	)
}