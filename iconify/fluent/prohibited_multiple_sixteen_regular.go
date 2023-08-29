package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedMultipleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 12A5 5 0 1 0 7 2a5 5 0 0 0 0 10Zm0-1a3.982 3.982 0 0 1-2.453-.84l5.613-5.613A4 4 0 0 1 7 11Zm2.453-7.16L3.84 9.453A4 4 0 0 1 9.453 3.84ZM13 7c0 .377-.035.745-.101 1.103a4 4 0 0 1-4.796 4.796a6.032 6.032 0 0 1-2.242-.007a5 5 0 0 0 7.031-7.031c.07.369.108.75.108 1.139Z"/>`),
		g.Group(children),
	)
}