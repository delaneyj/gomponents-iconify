package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyRideConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 13.8v7.3L39.6 24v10.2c0 .8-.7 1.5-1.6 1.5H6.5c-1.1 0-2-.7-2-1.5v-7.3L8.4 24V13.8c0-.8.7-1.5 1.6-1.5h31.5c1.1 0 2 .7 2 1.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.2 25.4c0-1.3 1.1-2.4 2.4-2.4s2.4 1.1 2.4 2.4v3.8M25.2 23v6.2m4.8-3.8c0-1.3 1.1-2.4 2.4-2.4s2.4 1.1 2.4 2.4v3.8m-21.5-9.5v9.5m0-2l4.3-4.3m-2.9 2.9l3.4 3.4m-1.7-8h11.2m-6.3 8v-8"/>`),
		g.Group(children),
	)
}