package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNpm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 8h22v7H11v2H7v-2H1zm6 0v7m7-7v7m3-4v4M4 11v4m7-4v1m9-1v4"/>`),
		g.Group(children),
	)
}