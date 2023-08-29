package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm6.113 5.654l-10.46 10.46a7.5 7.5 0 0 0 10.46-10.46ZM12 4.5a7.5 7.5 0 0 0-6.113 11.846l10.46-10.46A7.466 7.466 0 0 0 12 4.5Z"/>`),
		g.Group(children),
	)
}