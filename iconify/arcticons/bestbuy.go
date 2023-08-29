package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bestbuy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.22 11.24L4.5 18.86v10.42l7.5 7.48h31.5V11.24ZM9.63 21.75a1.94 1.94 0 1 1-1.93 1.94h0a1.93 1.93 0 0 1 1.93-1.94Z"/>`),
		g.Group(children),
	)
}