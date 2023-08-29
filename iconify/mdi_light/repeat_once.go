package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOnce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 7.5l-3.54 3.54l-.7-.71L18.09 8H6v4H5V7h13.09l-2.33-2.33l.7-.71L20 7.5M17 17v-4h1v5H4.91l2.33 2.33l-.7.71L3 17.5l3.54-3.54l.7.71L4.91 17H17m-7-3h1v-3h-1v-1h2v4h1v1h-3v-1Z"/>`),
		g.Group(children),
	)
}