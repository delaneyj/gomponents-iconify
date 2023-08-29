package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.03 17L12 20l-7.5-7.428A5 5 0 1 1 12 6.006a5 5 0 1 1 7.922 6.102M19 16v3m0 3v.01"/>`),
		g.Group(children),
	)
}