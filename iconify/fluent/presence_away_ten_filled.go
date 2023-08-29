package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 10A5 5 0 1 0 5 0a5 5 0 0 0 0 10Zm0-6.996v1.79l1.354 1.353a.5.5 0 1 1-.708.707l-1.5-1.5A.5.5 0 0 1 4 5V3.004a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}