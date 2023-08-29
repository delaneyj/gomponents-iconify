package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.22 16.78a.75.75 0 0 0 1.06 0L15.5 5.56v7.69a.75.75 0 0 0 1.5 0v-9.5a.75.75 0 0 0-.75-.75h-9.5a.75.75 0 0 0 0 1.5h7.69L3.22 15.72a.75.75 0 0 0 0 1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}