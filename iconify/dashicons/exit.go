package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3v2h2v10h-2v2h4V3h-4zm0 8V9H5.4l4.3-4.3l-1.4-1.4L1.6 10l6.7 6.7l1.4-1.4L5.4 11H13z"/>`),
		g.Group(children),
	)
}