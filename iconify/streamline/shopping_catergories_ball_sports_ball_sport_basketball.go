package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCatergoriesBallSportsBallSportBasketball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="6.5"/><path d="M7 .5v13m-4.9-2.23A5 5 0 0 0 4.5 7a5 5 0 0 0-2.4-4.27m9.8 0a5 5 0 0 0 0 8.54"/></g>`),
		g.Group(children),
	)
}