package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hikar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.657 10.5v-6h-6.5v6m0 13v20h6.5v-20M35 10.5H6v13h29l7-6.5l-7-6.5z"/>`),
		g.Group(children),
	)
}