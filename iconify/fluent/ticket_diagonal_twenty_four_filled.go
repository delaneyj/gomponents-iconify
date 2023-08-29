package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketDiagonalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.59 2.53a2.25 2.25 0 0 0-3.181 0L2.53 12.41a2.25 2.25 0 0 0 0 3.182l1.172 1.171c.51.511 1.227.42 1.66.162a1.25 1.25 0 0 1 1.713 1.713c-.257.433-.349 1.15.162 1.66L8.41 21.47a2.25 2.25 0 0 0 3.182 0l9.878-9.878a2.25 2.25 0 0 0 0-3.182l-1.171-1.172c-.51-.51-1.228-.42-1.661-.162a1.25 1.25 0 0 1-1.713-1.713c.258-.433.349-1.15-.162-1.66L15.591 2.53Z"/>`),
		g.Group(children),
	)
}