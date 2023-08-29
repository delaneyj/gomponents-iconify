package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoliceJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5.5a5 5 0 1 0 5 5a5 5 0 0 0-5-5zm2.057 2.182L5.5 4.739L3.443 2.682a3.442 3.442 0 0 1 4.114 0zm-4.875.761L4.739 5.5L2.682 7.557a3.442 3.442 0 0 1 0-4.114zm.761 4.875L5.5 6.261l2.057 2.057a3.442 3.442 0 0 1-4.114 0zm4.875-.761L6.261 5.5l2.057-2.057a3.442 3.442 0 0 1 0 4.114z" fill="currentColor"/>`),
		g.Group(children),
	)
}