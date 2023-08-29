package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAvailableTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 0a6 6 0 1 0 0 12A6 6 0 0 0 6 0ZM1.5 6a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Zm7.03-1.78a.75.75 0 0 1 0 1.06L5.78 8.03a.75.75 0 0 1-1.06 0l-1-1a.75.75 0 0 1 1.06-1.06l.47.47l2.22-2.22a.75.75 0 0 1 1.06 0Z"/>`),
		g.Group(children),
	)
}