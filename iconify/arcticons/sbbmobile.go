package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sbbmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.76 14.25L43.5 24l-9.75 9.75m-19.51 0L4.5 24l9.75-9.75M4.5 24h39M24 14.25v19.5"/>`),
		g.Group(children),
	)
}