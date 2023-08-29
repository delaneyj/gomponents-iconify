package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.338 6.713a2 2 0 0 1-.967-2.658L16.58 2.75L6.21 24.948l2.793 1.305c.468-1 1.658-1.434 2.66-.966s1.433 1.657.965 2.658l2.793 1.305L25.79 7.052l-2.794-1.305a2 2 0 0 1-2.658.966z"/>`),
		g.Group(children),
	)
}