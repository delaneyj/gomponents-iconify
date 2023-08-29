package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaClosedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c.4 0 .8.2.9.6L17.5 15H13v4c0 1.7-1.3 3-3 3s-3-1.3-3-3v-1h2v1c0 .6.4 1 1 1s1-.4 1-1v-4H6.5l4.6-12.4c.1-.4.5-.6.9-.6m0 3.9L9.4 13h5.3L12 5.9Z"/>`),
		g.Group(children),
	)
}