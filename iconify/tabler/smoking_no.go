package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmokingNo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 13v4m8-12v.5a2 2 0 0 0 2 2a2 2 0 0 1 2 2v.5M3 3l18 18m-4-8h3a1 1 0 0 1 1 1v2c0 .28-.115.533-.3.714M17 17H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h9"/>`),
		g.Group(children),
	)
}