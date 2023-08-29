package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareArrowForwardThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 3A4.5 4.5 0 0 0 3 7.5v17A4.5 4.5 0 0 0 7.5 29h8.775A9 9 0 1 1 29 16.275V7.5A4.5 4.5 0 0 0 24.5 3h-17ZM30 22.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-4.72-3.78a.75.75 0 1 0-1.06 1.06L25.44 21H22.5a4.25 4.25 0 0 0-4.25 4.25v.5a.75.75 0 0 0 1.5 0v-.5a2.75 2.75 0 0 1 2.75-2.75h2.94l-1.22 1.22a.75.75 0 1 0 1.06 1.06l2.5-2.5a.75.75 0 0 0 0-1.06l-2.5-2.5Z"/>`),
		g.Group(children),
	)
}