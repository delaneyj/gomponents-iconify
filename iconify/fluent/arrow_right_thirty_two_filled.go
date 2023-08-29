package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 16c0-.69.56-1.25 1.25-1.25h20.537l-7.432-7.628a1.25 1.25 0 1 1 1.79-1.744l9.497 9.747a1.246 1.246 0 0 1 0 1.75l-9.497 9.747a1.25 1.25 0 1 1-1.79-1.744l7.432-7.628H4.25C3.56 17.25 3 16.69 3 16Z"/>`),
		g.Group(children),
	)
}