package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noteless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.51 17.51V5.5H40.5a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2h-33a2 2 0 0 1-2-2V17.51ZM17.509 5.5L5.5 17.509m6 6.099h25m-13.991-9.001H36.5m-25 18.001h15m5.22 0h4.78"/>`),
		g.Group(children),
	)
}