package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 12A5 5 0 1 0 7 2a5 5 0 0 0 0 10Zm0-1.5a3.484 3.484 0 0 1-1.887-.552l4.835-4.835A3.5 3.5 0 0 1 7 10.5Zm1.887-6.448L4.052 8.887a3.5 3.5 0 0 1 4.835-4.835ZM13 7a6 6 0 0 1-7.14 5.892a5 5 0 0 0 7.031-7.031c.072.369.109.75.109 1.139Z"/>`),
		g.Group(children),
	)
}