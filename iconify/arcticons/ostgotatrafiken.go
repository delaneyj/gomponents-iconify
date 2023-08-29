package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ostgotatrafiken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.8 13.847a13.252 13.252 0 1 1-8.518-3.1m0 0H43.5m-28.267 0H4.5"/>`),
		g.Group(children),
	)
}