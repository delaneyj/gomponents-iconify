package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loopboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.72 12.5H16A11.5 11.5 0 0 0 4.5 24h0A11.5 11.5 0 0 0 16 35.5h16A11.5 11.5 0 0 0 43.5 24h0A11.5 11.5 0 0 0 32 12.5H21.881"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.701 5.68l-6.82 6.82l6.82 6.82"/>`),
		g.Group(children),
	)
}