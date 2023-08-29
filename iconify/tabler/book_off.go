package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19a9 9 0 0 1 9 0a9 9 0 0 1 5.899-1.096M3 6a9 9 0 0 1 2.114-.884m3.8-.21C9.984 5.076 11.03 5.44 12 6a9 9 0 0 1 9 0M3 6v13m9-13v2m0 4v7m9-13v11M3 3l18 18"/>`),
		g.Group(children),
	)
}