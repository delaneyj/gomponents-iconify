package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDashThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 3A4.5 4.5 0 0 0 3 7.5V11a1 1 0 1 0 2 0V7.5A2.5 2.5 0 0 1 7.5 5H11a1 1 0 1 0 0-2H7.5ZM21 3a1 1 0 1 0 0 2h3.5A2.5 2.5 0 0 1 27 7.5V11a1 1 0 1 0 2 0V7.5A4.5 4.5 0 0 0 24.5 3H21ZM5 21a1 1 0 1 0-2 0v3.5A4.5 4.5 0 0 0 7.5 29H11a1 1 0 1 0 0-2H7.5A2.5 2.5 0 0 1 5 24.5V21Zm24 0a1 1 0 1 0-2 0v3.5a2.5 2.5 0 0 1-2.5 2.5H21a1 1 0 1 0 0 2h3.5a4.5 4.5 0 0 0 4.5-4.5V21ZM9 15a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2H9Z"/>`),
		g.Group(children),
	)
}