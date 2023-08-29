package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 4a2 2 0 0 0-4 0v12h4V4zm5 4a2 2 0 0 0-4 0v8h4V8zM9 11a2 2 0 0 0-4 0v5h4v-5zm10 8H5a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/>`),
		g.Group(children),
	)
}