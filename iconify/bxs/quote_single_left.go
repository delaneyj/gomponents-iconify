package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteSingleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.203 6.98l.804-.161V4h-1c-2.784 0-4.906.771-6.309 2.292C6.81 8.34 7 10.97 7.006 11v8a1 1 0 0 0 1 1h7c1.103 0 2-.897 2-2v-7a1 1 0 0 0-1-1h-4.079c.022-.402.123-.912.429-1.396c.509-.801 1.466-1.347 2.847-1.624z"/>`),
		g.Group(children),
	)
}