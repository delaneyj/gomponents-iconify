package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresenceAwayTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 24c6.627 0 12-5.373 12-12S18.627 0 12 0S0 5.373 0 12s5.373 12 12 12Zm1-17.5v4.81l2.976 2.551a1.5 1.5 0 0 1-1.952 2.278l-3.5-3A1.5 1.5 0 0 1 10 12V6.5a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}