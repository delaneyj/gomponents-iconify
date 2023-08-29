package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandUpLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v9a1 1 0 1 0 2 0V6.414l7.293 7.293a1 1 0 0 0 1.414-1.414L6.414 5H13a1 1 0 0 0 1-1Zm10.5 1A2.5 2.5 0 0 1 27 7.5V16h-7.231a3.77 3.77 0 0 0-3.77 3.77V27H7.5A2.5 2.5 0 0 1 5 24.5V19a1 1 0 1 0-2 0v5.5A4.5 4.5 0 0 0 7.5 29h17a4.5 4.5 0 0 0 4.5-4.5v-17A4.5 4.5 0 0 0 24.5 3H19a1 1 0 1 0 0 2h5.5Z"/>`),
		g.Group(children),
	)
}