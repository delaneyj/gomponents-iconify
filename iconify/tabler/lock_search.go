package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 21H7a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h10m-9 0V7a4 4 0 1 1 8 0v4m-1 7a3 3 0 1 0 6 0a3 3 0 1 0-6 0m5.2 2.2L22 22"/>`),
		g.Group(children),
	)
}