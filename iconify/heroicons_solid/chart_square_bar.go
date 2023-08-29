package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartSquareBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5Zm9 4a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0V7Zm-3 2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0V9Zm-3 3a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0v-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}