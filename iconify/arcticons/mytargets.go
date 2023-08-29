package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mytargets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.51 21.51 0 0 0 2.5 24h0A21.51 21.51 0 0 0 24 45.5h0A21.51 21.51 0 0 0 45.5 24h0A21.51 21.51 0 0 0 24 2.5Zm0 4.26A17.24 17.24 0 1 1 6.76 24A17.24 17.24 0 0 1 24 6.76Zm0 4.31A12.93 12.93 0 0 0 11.07 24h0a12.93 12.93 0 0 0 25.86 0h0A12.93 12.93 0 0 0 24 11.07Zm0 4.36A8.57 8.57 0 1 1 15.43 24A8.57 8.57 0 0 1 24 15.43Zm0 4.31A4.26 4.26 0 0 0 19.74 24h0A4.26 4.26 0 0 0 24 28.26h0A4.26 4.26 0 0 0 28.26 24h0A4.26 4.26 0 0 0 24 19.74Z"/>`),
		g.Group(children),
	)
}