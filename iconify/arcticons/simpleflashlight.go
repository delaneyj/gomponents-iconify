package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simpleflashlight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm0 7.44A14.06 14.06 0 1 0 38.06 24h0A14.07 14.07 0 0 0 24 9.94Zm0 7.44A6.62 6.62 0 1 0 30.62 24A6.63 6.63 0 0 0 24 17.38Z"/>`),
		g.Group(children),
	)
}