package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimplemindLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5 2.5V24A21.5 21.5 0 0 0 24 45.5h0A21.5 21.5 0 0 0 45.5 24h0A21.5 21.5 0 0 0 24 2.5H2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.15 10.15V24A13.85 13.85 0 0 0 24 37.85h0A13.85 13.85 0 0 0 37.85 24h0A13.85 13.85 0 0 0 24 10.15H10.15Z"/>`),
		g.Group(children),
	)
}