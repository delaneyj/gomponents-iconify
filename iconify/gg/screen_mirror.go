package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 8h14v6h-3v2h5V6H3v10h5v-2H5V8Z"/><path d="M16.33 19L12 13l-4.33 6h8.66Z"/></g>`),
		g.Group(children),
	)
}