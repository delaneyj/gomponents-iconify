package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ytaxi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 7.45V24h37V7.45a2 2 0 0 0-2-1.95H7.45A2 2 0 0 0 5.5 7.45ZM24 24v18.5h16.55a2 2 0 0 0 2-2V24ZM5.5 24v16.55a2 2 0 0 0 2 2H24V24Z"/>`),
		g.Group(children),
	)
}