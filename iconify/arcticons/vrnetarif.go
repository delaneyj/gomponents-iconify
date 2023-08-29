package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vrnetarif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5a8.74 8.74 0 0 1-8.4-8.74v-5.68a8.74 8.74 0 0 1 8.74-8.75h0a8.74 8.74 0 0 1 8.74 8.75v2.84H15.6"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45.5A21.5 21.5 0 1 1 45.5 24A21.51 21.51 0 0 1 24 45.5Z"/>`),
		g.Group(children),
	)
}