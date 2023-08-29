package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAvailableTenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 10A5 5 0 1 0 5 0a5 5 0 0 0 0 10Zm2.104-5.896l-2.25 2.25a.5.5 0 0 1-.708 0l-1-1a.5.5 0 0 1 .708-.708l.646.647l1.896-1.897a.5.5 0 0 1 .708.708Z"/>`),
		g.Group(children),
	)
}