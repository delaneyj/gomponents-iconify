package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandWalmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8.04V3m3.5 7L20 7.5M15.5 14l4.5 2.5m-8-.54V21m-3.5-7L4 16.5M8.5 10L4 7.495"/>`),
		g.Group(children),
	)
}