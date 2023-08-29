package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxRemoveOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2M5 19v-2h3.1c.2.8.7 1.5 1.3 2m9.6 0h-4.4c.6-.5 1.1-1.2 1.3-2H19m0-2h-5v1c0 1.1-.9 2-2 2s-2-.9-2-2v-1H5V5h14v10m-4.9-8.5l1.4 1.4l-2.1 2.1l2.1 2.1l-1.4 1.4l-2.1-2.1l-2.1 2.1l-1.4-1.4l2.1-2.1l-2.1-2.1l1.4-1.4L12 8.6l2.1-2.1Z"/>`),
		g.Group(children),
	)
}