package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BritishAirways(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.178 22.7c-.911-1.1-.911-2.2-10.022-2.2H4c5.467 3.3 16.4 3.3 33.711 3.3c16.4 0 0 7.7 0 7.7l-6.378-1.1s7.29-2.2 9.111-4.4"/>`),
		g.Group(children),
	)
}