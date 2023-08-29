package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M0 16.7l3.2-3.8L0 9.1l1.5-1.3l3 3.5l3-3.5L9 9.1l-3.2 3.8L9 16.7L7.5 18l-3-3.6l-3 3.6L0 16.7m24 .2c0 .5-.4 1-1 1h-3c-1.1 0-2-.9-2-2v-2c0-1.1.9-2 2-2h2v-2h-4V8h5c.5 0 1 .4 1 1m-2 5h-2v2h2v-2m-6 2.9c0 .5-.4 1-1 1h-3c-1.1 0-2-.9-2-2v-6c0-1.1.9-2 2-2h2V5h2v11.9m-2-1v-6h-2v6h2z" fill="currentColor"/>`),
		g.Group(children),
	)
}