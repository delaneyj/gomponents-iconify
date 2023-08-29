package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoliceJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 1A6.5 6.5 0 1 0 14 7.5A6.5 6.5 0 0 0 7.5 1Zm2.6 2.84l-2.6 2.6l-2.6-2.6a4.443 4.443 0 0 1 5.2 0ZM3.84 4.9l2.6 2.6l-2.6 2.6a4.443 4.443 0 0 1 0-5.2Zm1.06 6.26l2.6-2.6l2.6 2.6a4.443 4.443 0 0 1-5.2 0Zm6.26-1.06l-2.6-2.6l2.6-2.6a4.443 4.443 0 0 1 0 5.2Z"/>`),
		g.Group(children),
	)
}