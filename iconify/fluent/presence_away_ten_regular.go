package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4.793V3.004a.5.5 0 0 0-1 0V5a.5.5 0 0 0 .146.354l1.5 1.5a.5.5 0 1 0 .708-.707L5 4.793ZM10 5A5 5 0 1 1 0 5a5 5 0 0 1 10 0ZM9 5a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/>`),
		g.Group(children),
	)
}