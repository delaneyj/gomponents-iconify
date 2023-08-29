package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pills(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8a5 5 0 1 0 10 0A5 5 0 1 0 3 8m10 9a4 4 0 1 0 8 0a4 4 0 1 0-8 0M4.5 4.5l7 7m8 3l-5 5"/>`),
		g.Group(children),
	)
}