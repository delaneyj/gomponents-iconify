package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AmericanFootballEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.53 2C2.47 2 1 5.5 1 5.5S2.47 9 5.53 9S10 5.5 10 5.5S8.6 2 5.53 2zM7 6H4a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1z" fill="currentColor"/>`),
		g.Group(children),
	)
}