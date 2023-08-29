package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21v-4m0-4v-1m0-7V3m-2 18h4M8 8v1h1m4 0h6l2-2l-2-2H9m5 9v3H6l-2-2l2-2h7M3 3l18 18"/>`),
		g.Group(children),
	)
}