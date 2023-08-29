package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unstoppable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.905 23.711a10.095 10.095 0 0 0 20.19 0v0l-.001-18.694a21.5 21.5 0 1 1-20.188 0h0l-.001 18.694Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.7 5.017v19.231h0a1.7 1.7 0 0 1-3.4 0h0V5.017"/>`),
		g.Group(children),
	)
}