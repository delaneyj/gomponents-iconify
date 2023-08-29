package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSortDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 28l-7-7l1.4-1.4l5.6 5.6l5.6-5.6L23 21z"/>`),
		g.Group(children),
	)
}