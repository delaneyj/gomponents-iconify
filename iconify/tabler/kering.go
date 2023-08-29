package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kering(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 15v-3.5a2.5 2.5 0 1 1 5 0V15m0-2h-5M3 9l3 6l3-6m0 11l6-16"/>`),
		g.Group(children),
	)
}