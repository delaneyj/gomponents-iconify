package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchError(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.097 5.904A6.5 6.5 0 0 0 4 10.504l.001 1h-2v-1a8.5 8.5 0 1 1 15.176 5.258l5.344 5.345l-1.414 1.414l-5.344-5.345A8.478 8.478 0 0 1 10.5 19h-1v-2h1a6.5 6.5 0 0 0 4.596-11.096ZM1.672 13.257L4.5 16.086l2.829-2.829l1.414 1.415L5.915 17.5l2.828 2.828l-1.414 1.415L4.5 18.914l-2.828 2.829l-1.414-1.415L3.086 17.5L.258 14.672l1.414-1.415Z"/>`),
		g.Group(children),
	)
}